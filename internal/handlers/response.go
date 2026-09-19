package handler

import (
	"encoding/json"
	"fmt"

)


type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var user User

func HandleRequest(req Request) string {
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
