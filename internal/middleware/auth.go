package middleware

import (
	"fmt"
	"strings"
	"github.com/aluprince/socket-server/internal/handlers"
)

const secretToken = "my-super-secret-token" //use os and your environment variable key in production

type Handler func(handler.Request) handler.Response


func AuthMiddleware(next Handler) Handler {
	return func(req handler.Request) handler.Response {
		auth := req.Headers.Get("Authorization")
		fmt.Printf(">>>Auth Token: %v", auth)

		parts := strings.SplitN(auth, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			return handler.Unauthorized()
		}

		if parts[1] != secretToken {
			return handler.Unauthorized()
		}

		return next(req)
	}
}

