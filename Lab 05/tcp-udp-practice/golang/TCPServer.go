package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	serverPort := 12000
	serverAddress := &net.TCPAddr{
		IP:   net.ParseIP("localhost"),
		Port: serverPort,
	}
	serverSocket, err := net.ListenTCP("tcp", serverAddress)
	if err != nil {
		panic(err)
	}
	defer serverSocket.Close()

	fmt.Println("The server is ready to receive")

	for {
		buffer := make([]byte, 2048)
		connection, err := serverSocket.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting TCP connection: ", err)
			continue
		}

		connection.Read(buffer)
		message := string(buffer)
		fmt.Printf("TCP receive %s FROM %v\n", message, connection.RemoteAddr())

		messageUpper := strings.ToUpper(message)

		if _, err := connection.Write([]byte(messageUpper)); err != nil {
			fmt.Println("Error sending TCP message: ", err)
			continue
		}
	}
}
