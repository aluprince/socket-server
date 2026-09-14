package server

import (
	"fmt"
	"log"
	"net"
	"github.com/aluprince/socket-server/internal/handlers"
)

var addr = "127.0.0.1:8000"


func RunServer() {
	fmt.Println(">>>Starting Server Hold On...")
	ln, err := net.Listen("tcp", addr)
	fmt.Printf("Listening at this address >>> %v >>>", addr)

	if err != nil {
		log.Fatalf("Error Found: %v", err)
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error Found: %v", err)
			continue
		}
		fmt.Println("New connection:", conn.RemoteAddr())

		buffer := make([]byte, 1024)
		
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatalf("Error; %v", err)
		}
		fmt.Println(string(buffer[:n]))
		requestParsed := handler.ParseRequest(buffer[:n])
		response := handler.HandleRequest(requestParsed)

		conn.Write([]byte(response))


		// conn.Write([]byte("HTTP/1.1 200 OK\r\n" +
        // "Content-Type: text/plain\r\n" +
        // "Content-Length: 12\r\n" +
        // "\r\n" +
        // "Hello World!"))

		conn.Close()
	}
}


