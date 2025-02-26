package main

import (
	"net"
)

func handleConnection(conn net.Conn) {

	d := make([]byte, 1024*1000)

	n, _ := conn.Read(d)

	conn.Write(d[:n])

	conn.Close()

}

func main() {
	listener, _ := net.Listen("tcp", ":7")

	defer listener.Close()

	for {

		conn, _ := listener.Accept()

		go handleConnection(conn)

	}

}
