from socket import *
serverPort = 12000
serverSocket = socket(AF_INET, SOCK_DGRAM)
serverSocket.bind(('0.0.0.0', serverPort))
print("The server is ready to receive")
while 1:
    message, clientAddress = serverSocket.recvfrom(2048)
    print("UDP receive", message, "FROM", clientAddress)
    modifiedMessage = message.upper()
    serverSocket.sendto(modifiedMessage, clientAddress)
