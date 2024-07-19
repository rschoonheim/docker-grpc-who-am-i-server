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

var whoamiServerAddr string
var tlsCertPath string
var tlsKeyPath string
var tlsCaPath string

func init() {
	whoamiServerAddr = os.Getenv("WHOAMI_SERVER_ADDR")
	tlsCertPath = os.Getenv("TLS_CERT_PATH")
	tlsKeyPath = os.Getenv("TLS_KEY_PATH")
	tlsCaPath = os.Getenv("CA_CERT_PATH")

	if whoamiServerAddr == "" {
		slog.Error("Missing environment variable", "variable", "WHOAMI_SERVER_ADDR")
		os.Exit(1)
	}

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

func makeTlsConfiguration() (*tls.Config, error) {
	backendCert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("failed to read certificate: %v", err)
	}

	backendKey, err := os.ReadFile(tlsKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key: %v", err)
	}

	cert, err := tls.X509KeyPair(backendCert, backendKey)
	if err != nil {
		log.Fatalf("failed to parse certificate: %v", err)
	}

	rootCa, err := os.ReadFile(tlsCaPath)
	if err != nil {
		slog.Error("Failed to read certificate", "error", err)
		os.Exit(1)
	}

	rootCaPool := x509.NewCertPool()
	if !rootCaPool.AppendCertsFromPEM(rootCa) {
		slog.Error("Failed to append certificate")
		os.Exit(1)
	}

	return &tls.Config{
		ServerName:   "client",
		Certificates: []tls.Certificate{cert},
		RootCAs:      rootCaPool,
	}, nil
}

func main() {
	tlsConfiguration, err := makeTlsConfiguration()
	if err != nil {
		slog.Error("Failed to create TLS configuration", "error", err)
		os.Exit(1)
	}
	credsClient := credentials.NewTLS(tlsConfiguration)

	conn, err := grpc.Dial(whoamiServerAddr, grpc.WithTransportCredentials(credsClient))
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
