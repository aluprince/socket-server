package handler

import "fmt"


// GET REQUESTS
func Health() Response {
    body := "Active Health!"

    responseString := "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/plain\r\n" +
        fmt.Sprintf("Content-Length: %d\r\n", len([]byte(body))) +
        "\r\n" +
        body

    response := Response{
        StatusCode: getStatusCode(responseString),
        Body: responseString,
    }

    return response
}

//POST REQUESTS

func userCreated() Response {
    body := "User Has Been Created"

    responseString := "HTTP/1.1 201 CREATED\r\n" +
        "Content-Type: text/plain\r\n" +
        fmt.Sprintf("Content-Lenght: %d\r\n", len([]byte(body))) +
        "\r\n" +
        body

    response := Response{
        StatusCode: getStatusCode(responseString),
        Body: responseString,
    }
    return response
}
