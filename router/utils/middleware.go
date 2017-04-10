package utils

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/tallduck/sailfish-backend/protobuf/auth"
)

// AuthenticationMiddleware checks for presence of token and sends it
// to be validated. If missing or invalid we return a 403 Forbidden to
// the response.
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if len(header) == 0 {
			UnauthorizedHTTP(w)
			return
		}

		header = strings.Replace(header, "Token ", "", 1)
		fmt.Println(header)

		reqBuffer := &auth.Request{
			Token: header,
		}

		fmt.Println(reqBuffer)

		conn, err := grpc.Dial("auth:8080", grpc.WithInsecure())
		if err != nil {
			grpclog.Fatalf("failed to dial: %v", err)
			ErrorHTTP(w)
			return
		}

		client := auth.NewAuthClient(conn)
		// Call auth microservice to validate token
		auth, err := client.Authenticate(context.Background(), reqBuffer)
		if err != nil {
			ErrorHTTP(w)
			return
		}

		if !auth.Status {
			UnauthorizedHTTP(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
