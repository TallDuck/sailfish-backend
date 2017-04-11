package grpc

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/tallduck/sailfish-backend/helpers"
)

// Dial ensures our connection to the gRPC server. It also encapsulates common
// options that we would normally set on every grpc.Dial
func Dial(destination string, currentRequest http.ResponseWriter) *grpc.ClientConn {
	conn, err := grpc.Dial(destination, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
		helpers.ErrorHTTP(currentRequest)
		return nil
	}

	return conn
}
