package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	serverPort := 12000
	serverAddress := &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: serverPort,
	}
	serverSocket, err := net.ListenUDP("udp", serverAddress)
	if err != nil {
		panic(err)
	}
	defer serverSocket.Close()

	fmt.Println("The server is ready to receive")

	for {
		buffer := make([]byte, 2048)
		n, clientAddress, err := serverSocket.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading UDP message: ", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("UDP receive %s FROM %v\n", message, clientAddress)

		messageUpper := strings.ToUpper(message)

		_, err = serverSocket.WriteToUDP([]byte(messageUpper), clientAddress)
		if err != nil {
			fmt.Println("Error sending UDP message: ", err)
			continue
		}
	}
}
