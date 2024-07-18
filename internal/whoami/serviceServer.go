package whoami

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WhoAmIServerImplementation - implementation of the WhoAmI gRPC service
type WhoAmIServerImplementation struct {
}

func (WhoAmIServerImplementation) GetWhoAmI(ctx context.Context, req *WhoAmIRequest) (*WhoAmIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWhoAmI not implemented")
}

func (WhoAmIServerImplementation) mustEmbedUnimplementedWhoAmIServer() {}
