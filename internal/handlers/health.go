package handler

import "fmt"


func Health() string {
    body := "Active Health!"

    return "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/plain\r\n" +
        fmt.Sprintf("Content-Length: %d\r\n", len([]byte(body))) +
        "\r\n" +
        body
}
