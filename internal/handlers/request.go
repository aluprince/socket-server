package handler

import (
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	Method string
	Path string
	Proto string
	Headers http.Header
	Body []byte
}


func ParseRequest(data []byte) Request {
    // turn raw bytes into structured HTTP request
	buffer := string(data)

	//Separating head from body
	parts := strings.Split(buffer, "\r\n\r\n")
	requestLine := strings.Split(parts[0], "\r\n")[0]
	requestParts := strings.Split(requestLine, " ")

	req := Request{
		Method: requestParts[0],
		Path: requestParts[1],
		Proto: requestParts[2],
	}
	fmt.Println("--------------Request Parsed----------------")
	fmt.Println("Method: ", req.Method)
	fmt.Println("Path: ", req.Path)
	fmt.Println("Protocol: ", req.Proto)

	
	return req
}