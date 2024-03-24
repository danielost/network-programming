import socket

def is_palindrome(num):
    return str(num) == str(num)[::-1]

def main():
    host = '0.0.0.0'
    port = 5000

    server = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    server.bind((host, port))

    print(f'Server is listening on {host}:{port}')

    while True:
        data, addr = server.recvfrom(1024)
        number = int.from_bytes(data, byteorder='big') 
        print(f'Received number: {number} from {addr}')

        if is_palindrome(number):
            server.sendto(b'YES', addr)
        else:
            server.sendto(b'NO', addr)

if __name__ == '__main__':
    main()
