from socket import *
serverName = "localhost"
serverPort = 12000
clientSocket = socket(AF_INET, SOCK_STREAM)
clientSocket.connect((serverName,serverPort))
numbers = bytes(input('Input two numbers separated by space:'), 'utf_8')
clientSocket.send(numbers)
sum = clientSocket.recv(1024)
print("From TCP Server:", sum.decode('utf_8'))
clientSocket.close()
