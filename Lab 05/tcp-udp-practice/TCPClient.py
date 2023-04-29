from socket import *
serverName = "localhost"
serverPort = 12000
clientSocket = socket(AF_INET, SOCK_STREAM)
clientSocket.connect((serverName,serverPort))
sentence = bytes(input('Input lowercase sentence: '), 'utf_8')
clientSocket.send(sentence)
modifiedSentence = clientSocket.recv(1024)
print("From TCP Server:", modifiedSentence.decode('utf_8'))
clientSocket.close()
