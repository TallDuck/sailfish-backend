package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/tallduck/sailfish-backend/auth/jwt"
	"github.com/tallduck/sailfish-backend/helpers"
	"github.com/tallduck/sailfish-backend/protobuf/auth"
)

type authServer struct {
}

func (s *authServer) Authenticate(ctx context.Context, request *auth.AuthRequest) (*auth.AuthResponse, error) {
	fmt.Println("Starting authentication request")

	if len(request.Username) > 0 && len(request.Password) > 0 {
		jwt := jwt.New()
		token, err := jwt.Generate(map[string]string{
			"username": request.Username,
		})

		if err != nil {
			return &auth.AuthResponse{Status: false}, err
		}

		return &auth.AuthResponse{Status: true, Token: token}, nil
	}

	return &auth.AuthResponse{Status: false}, nil
}

func (s *authServer) Validate(ctx context.Context, request *auth.TokenRequest) (*auth.TokenResponse, error) {
	fmt.Println("Starting token validation request")

	jwt := jwt.New()
	tokenValid, _ := jwt.Validate(request.Token)

	if tokenValid {
		return &auth.TokenResponse{Status: true}, nil
	}

	return &auth.TokenResponse{Status: false}, nil
}

func (s *authServer) Invalidate(ctx context.Context, request *auth.TokenRequest) (*auth.TokenResponse, error) {
	return &auth.TokenResponse{Status: false}, nil
}

func main() {
	port := helpers.GetEnv("APP_PORT", "8080")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthServer(grpcServer, new(authServer))

	fmt.Println(fmt.Sprintf("Listening on tcp %v", port))
	grpcServer.Serve(listen)
}
