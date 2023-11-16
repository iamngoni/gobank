package main

import "fmt"

func main() {
	fmt.Println("Yeah buddy")
	server := NewAPIServer{listenAddress: "8000"}
	server.run()
}
