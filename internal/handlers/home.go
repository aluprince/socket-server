//This package is for handling "/"
package handler


import (
    "fmt"
)


//GET REQUESTS
func Home() Response {
    body := "Hello World!" 

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


func NotFound() Response {
    body := "I'm Sorry We Couldn't Process Your Path!"

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

func BadRequest() Response {
     body := "Bad Request"

    responseString := "HTTP/1.1 400 OK\r\n" +
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

func Unauthorized() Response {
     body := "Unauthorized Request"

    responseString := "HTTP/1.1 400 OK\r\n" +
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
