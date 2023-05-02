package main

import (
	"fmt"
	"net"
)

func main() {
	serverPort := 12000
	serverAddress := &net.TCPAddr{
		IP:   net.ParseIP("localhost"),
		Port: serverPort,
	}
	serverSocket, err := net.DialTCP("tcp", nil, serverAddress)
	if err != nil {
		panic(err)
	}
	defer serverSocket.Close()

	fmt.Println("The client is ready to send")
	fmt.Printf("Enter message: ")
	var message string
	fmt.Scanln(&message)

	if _, err = serverSocket.Write([]byte(message)); err != nil {
		panic(err)
	}

	modifiedMessage := make([]byte, 2048)
	n, err := serverSocket.Read(modifiedMessage)
	fmt.Println("From TCP Server: ", string(modifiedMessage[:n]))
}
