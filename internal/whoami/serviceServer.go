package whoami

import (
	"context"
)

// WhoAmIServerImplementation - implementation of the WhoAmI gRPC service
type WhoAmIServerImplementation struct {
}

func (WhoAmIServerImplementation) GetWhoAmI(ctx context.Context, req *WhoAmIRequest) (*WhoAmIResponse, error) {

	return &WhoAmIResponse{
		Message: "Hello from the gRPC server",
	}, nil
}

func (WhoAmIServerImplementation) mustEmbedUnimplementedWhoAmIServer() {}
