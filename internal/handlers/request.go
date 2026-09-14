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
	lines := strings.Split(parts[0], "\r\n")
	requestLine := lines[0]
	requestParts := strings.Split(requestLine, " ")


	//Parsing HTTP Headers Cleanly
	Headers := make(map[string][]string)
	// headerSection := strings.Split(parts[0], "\r\n")[1]
	seperateHeaders := lines[1:]
	
	fmt.Printf(">>We have about %v headers >>", len(seperateHeaders))
	for i := 0; i < len(seperateHeaders); i++ {
		sep := strings.SplitN(seperateHeaders[i], ":", 2)

		key := strings.TrimSpace(sep[0])
		value := strings.TrimSpace(sep[1])

		Headers[key] = []string{value}
	}

	req := Request{
		Method: requestParts[0],
		Path: requestParts[1],
		Proto: requestParts[2],
		Headers: Headers,
	}

	fmt.Println("--------------Request Parsed----------------")
	fmt.Println("Method: ", req.Method)
	fmt.Println("Path: ", req.Path)
	fmt.Println("Protocol: ", req.Proto)

	fmt.Println("------------REQUEST HEADERS-------------")
	fmt.Println(">>>Headers: ", req.Headers)

	
	return req
}