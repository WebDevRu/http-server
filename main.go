package main

import (
	"fmt"
	"http-server/src/requests_listener"
)

func main() {
	requests_listener.Listen()
	fmt.Println("test")
}
