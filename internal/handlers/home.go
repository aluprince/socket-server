//This package is for handling "/"
package handler


import (
    "fmt"
)

func Home() string {
    body := "Hello World!" 

    return "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/plain\r\n" +
        fmt.Sprintf("Content-Length: %d\r\n", len([]byte(body))) +
        "\r\n" +
        body
}

func NotFound() string {
    body := "I'm Sorry We Couldn't Process Your Path!"

    return "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/plain\r\n" +
        fmt.Sprintf("Content-Length: %d\r\n", len([]byte(body))) +
        "\r\n" +
        body
}
