package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:6379") // start port listen:6379
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	accept, err := listen.Accept() // start port accept
	defer accept.Close()           // close port listen
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	// try to read accept data
	n, err := accept.Read(buf)
	if err != nil {
		fmt.Println("error reading from client: ", err.Error())
		os.Exit(1)
	}
	println(n)
	// accept successful , response a data
	accept.Write([]byte("PONG\r\n"))

}
