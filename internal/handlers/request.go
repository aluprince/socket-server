package handler

import (
	//"fmt"
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
	seperateHeaders := lines[1:]
	
	//fmt.Printf(">>We have about %v headers >>", len(seperateHeaders))
	for i := 0; i < len(seperateHeaders); i++ {
		sep := strings.SplitN(seperateHeaders[i], ":", 2)

		key := strings.TrimSpace(sep[0])
		value := strings.TrimSpace(sep[1])

		Headers[key] = []string{value}
	}

	//Handling the body of the request
	//fmt.Printf(">>Parts 1: %v", parts[1])
	bodyString := parts[1]


	req := Request{
		Method: requestParts[0],
		Path: requestParts[1],
		Proto: requestParts[2],
		Headers: Headers,
		Body: []byte(bodyString),
	}

	// fmt.Println("--------------Request Parsed----------------")
	// fmt.Println("Method: ", req.Method)
	// fmt.Println("Path: ", req.Path)
	// fmt.Println("Protocol: ", req.Proto)

	// fmt.Println("------------REQUEST HEADERS-------------")
	// fmt.Println(">>>Headers: ", req.Headers)

	// fmt.Println("---------------REQUEST BODY -----------------")
	// fmt.Println(">>>Body: ", req.Body)

	
	return req
}