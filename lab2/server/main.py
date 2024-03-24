import socket

def is_palindrome(num):
    return str(num) == str(num)[::-1]

def main():
    host = '0.0.0.0'
    port = 5001

    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.bind((host, port))
    server.listen(5)

    print(f'Server is listening on {host}:{port}')

    while True:
        client_socket, addr = server.accept()
        print(f'Client connected from {addr}')

        while True:
            data = client_socket.recv(1024)
            if not data:
                break

            number = int.from_bytes(data, byteorder='big') 
            print(f'Received number: {number}')

            if is_palindrome(number):
                client_socket.send(b'YES')
            else:
                client_socket.send(b'NO')

        client_socket.close()
        print('Client disconnected')

    server.close()

if __name__ == '__main__':
    main()
