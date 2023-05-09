package main

import (
	"fmt"
	"net"
)

func main() {
	serverPort := 12000
	serverAddress := &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: serverPort,
	}
	serverSocket, err := net.DialUDP("udp", nil, serverAddress)
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
	n, _, err := serverSocket.ReadFromUDP(modifiedMessage)
	if err != nil {
		panic(err)
	}
	fmt.Println("From UDP Server: ", string(modifiedMessage[:n]))
}
