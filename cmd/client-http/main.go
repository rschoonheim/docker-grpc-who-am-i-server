package main

import (
	"context"
	"docker-grpc-who-am-i-service/internal/whoami"
	"google.golang.org/grpc"
	"log/slog"
	"os"
)

var whoamiServerAddr string

func init() {
	whoamiServerAddr = os.Getenv("WHOAMI_SERVER_ADDR")
	if whoamiServerAddr == "" {
		slog.Error("Missing environment variable", "variable", "WHOAMI_SERVER_ADDR")
	}
}

func main() {
	slog.Info("Starting HTTP client", "whoami_server_addr", whoamiServerAddr)

	// grpc connection
	//
	conn, err := grpc.Dial(whoamiServerAddr, grpc.WithInsecure())
	if err != nil {
		slog.Error("Failed to connect to gRPC server", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()

	client := whoami.NewWhoAmIClient(conn)

	// call the gRPC service
	//
	whoamiResponse, err := client.GetWhoAmI(ctx, &whoami.WhoAmIRequest{})
	if err != nil {
		slog.Error("Failed to call WhoAmI", "error", err)
		os.Exit(1)
	}

	slog.Info("GetWhoAMI response", "response", whoamiResponse.String())
}
