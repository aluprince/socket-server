package middleware

import (
	"log"
	"github.com/aluprince/socket-server/internal/handlers"
	"time"
)



func LogMiddleware(next Handler) Handler {
	return func(req handler.Request) handler.Response {

		start := time.Now()
		response := next(req)
		latency := time.Since(start)

		log.Printf("[REQUEST]: %v | %v | latency=%v", req.Method, req.Path, latency)

		return response
	}
}

func LogResponseMiddleware(next Handler) Handler {
    return func(req handler.Request) handler.Response {

		start := time.Now()
        resp := next(req)
		latency := time.Since(start)
        log.Printf("[RESPONSE] %s %s => %s | RESPONSE LATENCY>> %s ", req.Method, req.Path, resp.StatusCode, latency)
        log.Printf("[BODY] %s", resp.Body)
        return resp
    }
}
