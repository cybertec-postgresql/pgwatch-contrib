package main

import (
	"context"
	"encoding/json"
	"log"
	"sync" // <--- added for mutex

	"github.com/cybertec-postgresql/pgwatch3/rpc/proto"
	"github.com/destrex271/pgwatch3_rpc_server/sinks"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KafkaProdReceiver struct {
	mu           sync.RWMutex           // <--- added mutex to protect connRegistry
	connRegistry map[string]*kafka.Conn // <--- renamed from conn_regisrty
	uri          string
	auto_add     bool
	sinks.SyncMetricHandler
}

func (r *KafkaProdReceiver) HandleSyncMetric() {
	for {
		req, ok := r.GetSyncChannelContent()
		if !ok {
			return // channel closed
		}

		var err error
		switch req.Operation {
		case pb.SyncOperation_ADD:
			err = r.AddTopicIfNotExists(req.GetDBName()) // <--- mutex used inside AddTopicIfNotExists
		case pb.SyncOperation_DELETE:
			err = r.CloseConnectionForDB(req.GetDBName()) // <--- mutex used inside CloseConnectionForDB
		}

		if err != nil {
			log.Printf("[ERROR] error handling Kafka SyncMetric operation: %s", err)
		}
	}
}

func NewKafkaProducer(host string, topics []string, partitions []int, auto_add bool) (kpr *KafkaProdReceiver, err error) {
	connRegistry := make(map[string]*kafka.Conn)
	partitions_len := len(partitions)
	for index, topic := range topics {
		var conn *kafka.Conn
		if partitions_len > 0 {
			conn, err = kafka.DialLeader(context.Background(), "tcp", host, topic, partitions[index])
		} else {
			conn, err = kafka.DialLeader(context.Background(), "tcp", host, topic, 0)
		}
		if err != nil {
			return
		}

		connRegistry[topic] = conn
	}
	kpr = &KafkaProdReceiver{
		connRegistry:      connRegistry,
		uri:               host,
		SyncMetricHandler: sinks.NewSyncMetricHandler(1024),
		auto_add:          auto_add,
	}

	go kpr.HandleSyncMetric() // start background sync goroutine
	return kpr, nil
}

func (r *KafkaProdReceiver) AddTopicIfNotExists(dbName string) error {
	r.mu.Lock()         // <--- lock for write
	defer r.mu.Unlock() // <--- unlock automatically at function exit

	if _, ok := r.connRegistry[dbName]; ok {
		return nil
	}

	new_conn, err := kafka.DialLeader(context.Background(), "tcp", r.uri, dbName, 0)
	if err != nil {
		return err
	}

	r.connRegistry[dbName] = new_conn
	log.Println("[INFO]: Added Database " + dbName + " to sink")
	return nil
}

func (r *KafkaProdReceiver) CloseConnectionForDB(dbName string) error {
	r.mu.Lock() // <--- lock for write
	conn, ok := r.connRegistry[dbName]
	if ok {
		delete(r.connRegistry, dbName)
	}
	r.mu.Unlock() // <--- unlock here

	if !ok {
		return nil
	}

	err := conn.Close()
	if err != nil {
		return err
	}

	log.Println("[INFO]: Deleted Database " + dbName + " from sink")
	return nil
}

func (r *KafkaProdReceiver) UpdateMeasurements(ctx context.Context, msg *pb.MeasurementEnvelope) (*pb.Reply, error) {
	DBName := msg.GetDBName()

	// Reading the map without lock - safe because writes are rare
	conn, ok := func() (*kafka.Conn, bool) {
		r.mu.RLock() // optional read lock
		defer r.mu.RUnlock()
		conn, ok := r.connRegistry[DBName]
		return conn, ok
	}()

	if !ok {
		log.Println("[WARNING]: Connection does not exist for database " + DBName)
		if r.auto_add {
			log.Println("[INFO]: Adding database " + DBName + " since Auto Add is enabled")
			err := r.AddTopicIfNotExists(DBName) // safe with write lock
			if err != nil {
				log.Println("[ERROR]: Unable to create new connection")
				return nil, err
			}
			// read again after adding
			r.mu.RLock()
			conn = r.connRegistry[DBName]
			r.mu.RUnlock()
		} else {
			return nil, status.Error(codes.FailedPrecondition, "auto add not enabled")
		}
	}

	json_data, err := json.Marshal(msg)
	if err != nil {
		log.Println("Unable to convert measurements data to json")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = conn.WriteMessages(kafka.Message{Value: json_data})
	if err != nil {
		log.Println("Failed to write messages!")
		return nil, err
	}

	log.Println("[INFO]: Measurements Written to topic - ", DBName)
	return &pb.Reply{}, nil
}
