import socketserver as ss
import sys


def main():
    port = int(sys.argv[1] if len(sys.argv) > 1 else 9090)
    msg = sys.argv[2] if len(sys.argv) > 2 else "Hello, world !"
    data = f"HTTP/1.1 200 OK\r\nContent-Length: {len(msg)}\r\nContent-Type: text/plain; charset=utf-8\r\n\r\n{msg}"
    http_resp = data.encode()
    addr = ("127.0.0.1", port)
    print(f"started listening on {addr}")
    ss.TCPServer(addr, lambda *a: [a[0].sendall(http_resp), a[0].close()]).serve_forever()


if __name__ == "__main__":
    main()
