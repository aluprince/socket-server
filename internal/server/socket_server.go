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
		log.Printf("Error Found: %v", err)
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error Found: %v", err)
			continue
		}
		fmt.Println("New connection:", conn.RemoteAddr())
		go handleConnection(conn)

	}
}



func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
		
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error; %v", err)
	}
	fmt.Println(string(buffer[:n]))
	requestParsed := handler.ParseRequest(buffer[:n])
	response := handler.HandleRequest(requestParsed)

	fmt.Printf("RESPONSE: %q\n", response)

	k, err := conn.Write([]byte(response))
	if err != nil {
    	log.Printf("Write error: %v", err)
	}

	fmt.Printf("Wrote %d bytes\n", k)
}
