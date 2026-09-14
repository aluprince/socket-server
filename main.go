package main

import (
	"fmt"
	"github.com/aluprince/socket-server/internal/server"
)


func main(){
	fmt.Println(">>>Running Server From Main >>>")
	server.RunServer()
	// if err != nil {
	// 	fmt.Println("Testing Error: ", err)
	// }
	// fmt.Println("Server: ", s)
}