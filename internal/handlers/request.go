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
	
	
	for i := 0; i < len(seperateHeaders); i++ {
		sep := strings.SplitN(seperateHeaders[i], ":", 2)

		key := strings.TrimSpace(sep[0])
		value := strings.TrimSpace(sep[1])

		Headers[key] = []string{value}
	}

	bodyString := parts[1]

	req := Request{
		Method: requestParts[0],
		Path: requestParts[1],
		Proto: requestParts[2],
		Headers: Headers,
		Body: []byte(bodyString),
	}

	
	return req
}