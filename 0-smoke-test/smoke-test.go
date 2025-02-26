package main

import (
	"net"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	d := make([]byte, 1024*1000)

	n, e := conn.Read(d)

	if e != nil {
		return
	}

	conn.Write(d[:n])

}

func main() {
	listener, _ := net.Listen("tcp", ":7")

	defer listener.Close()

	for {

		conn, _ := listener.Accept()

		go handleConnection(conn)

	}

}
