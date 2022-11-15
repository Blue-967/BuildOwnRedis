package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:6379") // start port listen:6379
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	// close port listen

	for {
		conn, err := listen.Accept() // start port accept
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		// try to read accept data
		if _, err := conn.Read(buf); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("error reading from client: ", err.Error())
				os.Exit(1)
			}
		}
		// accept successful , response a data
		conn.Write([]byte("+PONG\r\n"))
	}
}
