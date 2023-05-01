from socket import *
serverPort = 12000
serverSocket = socket(AF_INET, SOCK_DGRAM)
serverSocket.bind(('0.0.0.0', serverPort))
print("The server is ready to receive")
while 1:
    message, clientAddress = serverSocket.recvfrom(2048)
    print(f"UDP receive {message.decode('utf_8')} FROM {clientAddress}")
    a,b = message.decode('utf_8').split(' ')
    sum = int(a) + int(b)
    serverSocket.sendto(str(sum).encode('utf_8'), clientAddress)
