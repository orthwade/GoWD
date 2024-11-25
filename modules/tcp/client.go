package tcp

import (
	"fmt"
	"net"
	"os"
)

func StartClient(address string) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		os.Exit(1)
	}

	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error while receiving data:", err)
		os.Exit(1)
	}
	fmt.Println("Received from server: ", string(buffer[:n]))

}
