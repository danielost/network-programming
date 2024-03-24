import socket

def main():
    host = '172.17.0.2'
    port = 5000

    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect((host, port))
    print('Connected to server')

    number = 54345
    client.send(number.to_bytes(4, byteorder='big'))

    data = client.recv(1024)
    print(f'Received from server: {data.decode()}')

    number = 54344
    client.send(number.to_bytes(4, byteorder='big'))

    data = client.recv(1024)
    print(f'Received from server: {data.decode()}')

    client.close()
    print('Connection closed')

if __name__ == '__main__':
    main()
