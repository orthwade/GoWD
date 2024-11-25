package tcp

import (
	"fmt"
	"net"
	"os"
)

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error while creating listener: ", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error while acceting connection", err)
			continue
		}

		defer conn.Close()

		fmt.Println("Connection from", conn.RemoteAddr())

		_, err = conn.Write([]byte("Hello from the server!\n"))

		if err != nil {
			fmt.Println("Error while sending data:", err)
			continue
		}
	}
}
