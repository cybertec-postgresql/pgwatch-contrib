// Command gen copies pgwatch.proto from the pinned pgwatch dependency into
// this package and runs protoc against that copy, so the generated Go/gRPC
// bindings always match the upstream schema. Invoked via `go generate` in
// ../doc.go.
package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	// Blank-imported so `go mod tidy` keeps this pinned in go.mod/go.sum;
	// its on-disk module directory is resolved below to source pgwatch.proto.
	_ "github.com/cybertec-postgresql/pgwatch/v6/api/pb"
)

const upstreamModule = "github.com/cybertec-postgresql/pgwatch/v6"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "gen:", err)
		os.Exit(1)
	}
}

func run() error {
	upstreamDir, err := moduleDir(upstreamModule)
	if err != nil {
		return fmt.Errorf("resolve %s: %w", upstreamModule, err)
	}

	upstreamProto := filepath.Join(upstreamDir, "api", "pb", "pgwatch.proto")
	if _, err := os.Stat(upstreamProto); err != nil {
		return fmt.Errorf("upstream pgwatch.proto not found at %s: %w", upstreamProto, err)
	}

	if err := copyFile(upstreamProto, "pgwatch.proto"); err != nil {
		return fmt.Errorf("copy pgwatch.proto from upstream: %w", err)
	}

	cmd := exec.Command("protoc",
		"--go_out=.", "--go_opt=paths=source_relative",
		"--go-grpc_out=.", "--go-grpc_opt=paths=source_relative",
		"pgwatch.proto",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("protoc: %w", err)
	}
	return nil
}

func moduleDir(mod string) (string, error) {
	out, err := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", mod).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() { _ = in.Close() }()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() { _ = out.Close() }()

	_, err = io.Copy(out, in)
	return err
}
