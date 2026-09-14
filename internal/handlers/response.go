package handler

import (
	"fmt"
)


func HandleRequest(req Request) string {
	fmt.Println(">>> Testing Request Handler")
	if req.Path == "/" {
		return Home()
	}

	if req.Path == "/health" {
		return Health()
	}

	return NotFound()
}
