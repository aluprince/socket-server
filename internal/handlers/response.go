package handler

import (
	"encoding/json"
	"fmt"
	"strings"

)

type Response struct {
    StatusCode string
    Body       string
}

type Handler func(Request) Response

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var user User


func HandleRequest(req Request) Response {
	fmt.Println(">>> Testing Request Handler")
	if req.Path == "/" {
		return Home()
	}
	
	if req.Path == "/health" {
		return Health()
	}


	if req.Path == "/users" && req.Method == "POST" {
		if req.Headers.Get("Content-Type") == "application/json" {
			err := json.Unmarshal(req.Body, &user)

			if err != nil {
				return BadRequest()
			}

			return userCreated()			
		}
		return BadRequest()
	}

	
	return NotFound()
}


func getStatusCode(response string) string {
	lines := strings.SplitN(response, "\r\n", 2)

	if len(lines) == 0 {
		return "Unknown"
	}
	return lines[0]
}