package main

import (
	"docker-grpc-who-am-i-service/internal/whoami"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	whoami.RegisterWhoAmIServer(
		grpcServer,
		whoami.WhoAmIServerImplementation{},
	)

	slog.Info("gRPC server started", "port", 8080)

	grpcServer.Serve(listener)

}
