package handler

import "fmt"

// GET REQUESTS
func Health() string {
    body := "Active Health!"

    return "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/plain\r\n" +
        fmt.Sprintf("Content-Length: %d\r\n", len([]byte(body))) +
        "\r\n" +
        body
}

//POST REQUESTS

func userCreated() string {
    body := "User Has Been Created"

    return "HTTP/1.1 201 OK\r\n" +
        "Content-Type: text/plain\r\n" +
        fmt.Sprintf("Content-Lenght: %d\r\n", len([]byte(body))) +
        "\r\n" +
        body
}
