from socket import *
serverName = "localhost"
serverPort = 12000
clientSocket = socket(AF_INET, SOCK_DGRAM)
message = bytes(input("Input lowercase sentence:"), 'utf_8')
clientSocket.sendto(message,(serverName, serverPort))
modifiedMessage, serverAddress = clientSocket.recvfrom(2048)
print("From UDP Server:", modifiedMessage.decode('utf_8'))
clientSocket.close()
