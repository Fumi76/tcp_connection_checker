package main

import (
    "log"
    "net"
)

func main() {
    connection, error := net.Dial("tcp", "localhost:50051");

    if error != nil {
        log.Printf("ERROR %s", error);
        return
    }
    defer connection.Close()
    log.Printf("Success")
}
