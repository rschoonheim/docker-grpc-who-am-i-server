package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"docker-grpc-who-am-i-service/internal/whoami"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"log/slog"
	"os"
)

var tlsCertPath string
var tlsKeyPath string
var tlsCaPath string

func init() {
	tlsCertPath = os.Getenv("TLS_CERT_PATH")
	tlsKeyPath = os.Getenv("TLS_KEY_PATH")
	tlsCaPath = os.Getenv("CA_CERT_PATH")

	if tlsCertPath == "" {
		slog.Error("Missing environment variable", "variable", "TLS_CERT_PATH")
		os.Exit(1)
	}

	if tlsKeyPath == "" {
		slog.Error("Missing environment variable", "variable", "TLS_KEY_PATH")
		os.Exit(1)
	}

	if tlsCaPath == "" {
		slog.Error("Missing environment variable", "variable", "CA_CERT_PATH")
		os.Exit(1)
	}
}

func main() {

	tlsCert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		slog.Error("Failed to read certificate", "error", err)
		os.Exit(1)
	}

	rootCa, err := os.ReadFile(tlsCaPath)
	if err != nil {
		slog.Error("Failed to read certificate", "error", err)
		os.Exit(1)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(rootCa) {
		slog.Error("Failed to append certificate")
		os.Exit(1)
	}

	credsClient := credentials.NewTLS(&tls.Config{
		ServerName: "client",
		RootCAs:    pool,
		Certificates: []tls.Certificate{{
			Certificate: [][]byte{tlsCert},
		}},
	})

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(credsClient))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := whoami.NewWhoAmIClient(conn)

	whoamiResponse, err := client.GetWhoAmI(ctx, &whoami.WhoAmIRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %s", whoamiResponse.Message)

}
