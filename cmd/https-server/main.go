package main

import (
	"crypto/tls"
	"crypto/x509"
	"docker-grpc-who-am-i-service/internal/whoami"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"log/slog"
	"net"
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

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(rootCa) {
		slog.Error("Failed to append certificate")
		os.Exit(1)
	}

	return &tls.Config{
		ServerName:   "server",
		RootCAs:      pool,
		Certificates: []tls.Certificate{cert},
	}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	tlsConfiguration, err := makeTlsConfiguration()
	if err != nil {
		log.Fatalf("failed to create TLS configuration: %v", err)
		os.Exit(1)
	}
	transportCredentials := credentials.NewTLS(tlsConfiguration)

	grpcServer := grpc.NewServer(
		grpc.Creds(transportCredentials),
	)

	whoami.RegisterWhoAmIServer(
		grpcServer,
		whoami.WhoAmIServerImplementation{},
	)

	slog.Info("gRPC server started", "port", 8080)

	grpcServer.Serve(listener)
}
