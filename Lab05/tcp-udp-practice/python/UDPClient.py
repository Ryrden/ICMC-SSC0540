from socket import *
serverName = "localhost"
serverPort = 12000
clientSocket = socket(AF_INET, SOCK_DGRAM)
message = bytes(input("Input two numbers separated by space:"), 'utf_8')
clientSocket.sendto(message,(serverName, serverPort))
modifiedMessage, serverAddress = clientSocket.recvfrom(2048)
print(f"From UDP Server: {modifiedMessage.decode('utf_8')}")
clientSocket.close()
