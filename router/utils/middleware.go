package utils

import (
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"github.com/tallduck/sailfish-backend/helpers"
	"github.com/tallduck/sailfish-backend/protobuf/auth"
	"github.com/tallduck/sailfish-backend/router/grpc"
)

// AuthenticationMiddleware checks for presence of token and sends it
// to be validated. If missing or invalid we return a 403 Forbidden to
// the response.
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if len(header) == 0 {
			helpers.UnauthorizedHTTP(w)
			return
		}

		header = strings.Replace(header, "Token ", "", 1)
		reqBuffer := &auth.TokenRequest{
			Token: header,
		}

		conn := grpc.Dial("auth:8080", w)
		if conn == nil {
			return
		}

		// Call auth microservice to validate token
		client := auth.NewAuthClient(conn)
		auth, err := client.Validate(context.Background(), reqBuffer)
		if err != nil {
			helpers.ErrorHTTP(w)
			return
		}

		if !auth.Status {
			helpers.UnauthorizedHTTP(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
