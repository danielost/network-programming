import socket

def main():
    host = '127.0.0.1'
    port = 5000

    client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    number = 12322
    number_bytes = number.to_bytes(4, byteorder='big')

    client.sendto(number_bytes, (host, port))
    print(f'Sent number {number} to server')

    response, addr = client.recvfrom(1024)
    print(f'Received from server: {response.decode()}')

    client.close()

if __name__ == '__main__':
    main()
