# Socket Server

A small HTTP server built from scratch in Go using raw TCP connections.

The goal of this project was to understand what happens underneath frameworks such as `net/http`: accepting TCP connections, reading raw bytes, parsing HTTP requests, constructing HTTP responses, handling concurrency, and building middleware.

This is an educational project, not a production-ready HTTP implementation.

## Features

* Raw TCP server using Go's `net` package
* Manual HTTP request parsing
* HTTP request line and header parsing
* Request body handling
* JSON request parsing
* Basic routing
* HTTP response construction
* HTTP status codes
* Concurrent connection handling with goroutines
* Authentication middleware
* Request logging middleware
* Response logging
* Handler latency measurement
* Basic error handling

## Architecture

The server follows this general flow:

```text
Client
  │
  │ TCP connection
  ▼
net.Listener
  │
  ▼
Accept connection
  │
  ▼
goroutine
  │
  ▼
Read raw bytes
  │
  ▼
Parse HTTP request
  │
  ▼
Middleware
  │
  ├── Logging
  │
  └── Authentication
  │
  ▼
Request Handler
  │
  ▼
Build HTTP response
  │
  ▼
Write response
  │
  ▼
Close connection
```

## Project Structure

```text
socket-server/
├── go.mod
├── main.go
├── README.md
└── internal/
    ├── handlers/
    │   ├── health.go
    │   ├── home.go
    │   └── request.go
    ├── middleware/
    │   └── logging.go
    └── server/
        └── socket_server.go
```

## Running the Server

Clone the repository:

```bash
git clone https://github.com/aluprince/socket-server.git
cd socket-server
```

Run the server:

```bash
go run main.go
```

The server listens on:

```text
127.0.0.1:8000
```

## Endpoints

### GET /

Returns the home response.

```bash
curl http://127.0.0.1:8000/
```

### GET /health

Returns the server health response.

```bash
curl http://127.0.0.1:8000/health
```

### POST /users

Accepts a JSON payload and requires a Bearer token.

```bash
curl -X POST http://127.0.0.1:8000/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer my-super-secret-token" \
  -d '{"name":"billy","age":88}'
```

Example response:

```text
HTTP/1.1 201 Created
Content-Type: text/plain
Content-Length: 21

User Has Been Created
```

## Middleware

The server uses middleware to wrap request handlers.

The basic model is:

```text
Request
   │
   ▼
Logging Middleware
   │
   ▼
Authentication Middleware
   │
   ▼
Request Handler
   │
   ▼
Response
```

Middleware receives the next handler and returns another handler:

```go
func Middleware(next Handler) Handler
```

This makes it possible to add functionality around handlers without putting that functionality directly into every route.

The project currently uses middleware for:

* Request logging
* Authentication
* Response logging
* Latency measurement

## Concurrency

Each accepted TCP connection is handled in its own goroutine:

```go
go handleConnection(conn)
```

This allows multiple clients to be processed concurrently rather than forcing every connection to wait for the previous one to finish.

Conceptually:

```text
             TCP Listener
                  │
       ┌──────────┼──────────┐
       ▼          ▼          ▼
  Goroutine A Goroutine B Goroutine C
       │          │          │
    Request A  Request B  Request C
```

## HTTP Parsing

One of the main purposes of this project was understanding the structure of an HTTP request at the byte level.

A request generally consists of:

```text
Request Line
Headers
Blank Line
Body
```

For example:

```text
POST /users HTTP/1.1
Host: 127.0.0.1:8000
Content-Type: application/json
Content-Length: 28

{"name":"billy","age":88}
```

The server manually separates these components before passing the parsed request to the handler.

HTTP headers are separated using CRLF:

```text
\r\n
```

and the header section is separated from the body using:

```text
\r\n\r\n
```

## Latency

The logging middleware measures the time spent processing a request through the application handler:

```go
start := time.Now()

response := next(req)

latency := time.Since(start)
```

This measures application-side processing time rather than the complete network round-trip between the client and server.

Example:

```text
[REQUEST]: POST | /users | latency=31.7µs
```

This distinction became an important part of understanding the difference between application latency and network latency.

## Authentication

The `/users` endpoint uses a simple Bearer token for educational purposes.

Example:

```text
Authorization: Bearer my-super-secret-token
```

The middleware validates the token before allowing the request to reach the handler.

The hardcoded token is intentionally simple and should **not** be used for a real production application.

## What I Learned

This project was primarily about understanding backend systems below the framework level.

Key concepts explored:

* TCP sockets
* Network connections
* HTTP request structure
* HTTP headers
* HTTP request bodies
* JSON
* HTTP status codes
* Concurrent programming with goroutines
* Middleware and function composition
* Authentication
* Logging
* Latency measurement
* Separation between transport, middleware, and application logic

Building the server manually also demonstrated what higher-level HTTP frameworks abstract away.

## Limitations

This server is intentionally simplified.

It does **not** attempt to implement a production-grade HTTP server.

Notable limitations include:

* Assumes the HTTP request can be read in a single TCP read
* Fixed-size request buffer
* No persistent connection handling
* No TLS
* No HTTP/2
* No sophisticated routing
* No database
* Hardcoded authentication token
* Simplified HTTP parsing
* Simplified error handling

These limitations are intentional because the purpose of the project was learning rather than replacing Go's production HTTP stack.

For production applications, Go's standard `net/http` package should be used instead of this implementation.

## Why I Built This

Rather than immediately relying on a framework, I wanted to understand what happens between:

```text
TCP connection
      ↓
HTTP bytes
      ↓
HTTP request
      ↓
Middleware
      ↓
Handler
      ↓
HTTP response
      ↓
TCP connection
```

The project served as a practical introduction to the abstractions used by modern backend frameworks.
