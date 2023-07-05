from socket import *
serverPort = 12000
serverSocket = socket(AF_INET,SOCK_STREAM)
serverSocket.bind(('0.0.0.0',serverPort))
serverSocket.listen(1)
print('The server is ready to receive')
while 1:
     connectionSocket, addr = serverSocket.accept()
     numbers = connectionSocket.recv(1024)
     print(f"TCP Receive {numbers.decode('utf_8')} FROM {addr}")
     a,b = numbers.decode('utf_8').split(' ')
     sum = int(a) + int(b)
     connectionSocket.send(str(sum).encode('utf_8'))
     connectionSocket.close()
