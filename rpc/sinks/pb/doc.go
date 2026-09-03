// Package pb contains the Go/gRPC bindings for pgwatch's RPC sink protocol.
//
// pgwatch.proto is not hand-maintained here: it is copied on every
// `go generate` from the pgwatch.proto shipped by the pinned
// github.com/cybertec-postgresql/pgwatch/v6 dependency (see api/pb in that
// module), so the schema never drifts from upstream. To pick up a schema
// change, bump that dependency (`go get -u github.com/cybertec-postgresql/pgwatch/v6`)
// and re-run go generate.
//
// To generate the Go code from the protobuf definitions, you need to install:
//   - protoc (Protocol Buffers compiler)
//   - protoc-gen-go (Go plugin for protoc)
//   - protoc-gen-go-grpc (Go gRPC plugin for protoc)
//
// On Windows:
//
//	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//	winget install protobuf
//
// On Linux/macOS:
//
//	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//	# Install protoc via package manager (apt, brew, etc.)
//
// Then run: go generate ./sinks/pb/
package pb

// Fetch pgwatch.proto from the pinned pgwatch dependency and generate protobuf files
//go:generate go run ./gen
