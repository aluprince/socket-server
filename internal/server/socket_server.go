package server

import (
	"fmt"
	"log"
	"net"
	"github.com/aluprince/socket-server/internal/handlers"
	"github.com/aluprince/socket-server/internal/middleware"
)

var addr = "127.0.0.1:8000"


func RunServer() {
	log.Printf(">>>Starting Server Hold On...")
	ln, err := net.Listen("tcp", addr)
	fmt.Printf("Listening at this address >>> %v >>>", addr)

	if err != nil {
		log.Printf("Error Found: %v", err)
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error Found: %v", err)
			continue
		}
		log.Printf("New connection: %v", conn.RemoteAddr())
		go handleConnection(conn)

	}
}



func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
		
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error; %v", err)
	}
	fmt.Println(string(buffer[:n]))

	requestParsed := handler.ParseRequest(buffer[:n])
	
	appHandler := middleware.LogResponseMiddleware(middleware.LogMiddleware(middleware.AuthMiddleware(handler.HandleRequest)))

	response := appHandler(requestParsed)


	k, err := conn.Write([]byte(response.Body))
	if err != nil {
    	log.Printf("Write error: %v", err)
	}


 	log.Printf("Wrote %d bytes\n", k)
}
