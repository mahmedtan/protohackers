package main

import (
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	d := make([]byte, 1024*1000)

	n, err := conn.Read(d)

	if err != nil {
		if err != io.EOF {
			fmt.Println("Read error:", err)
		}
		return
	}

	_, err2 := conn.Write(d[:n])

	if err2 != nil {
		fmt.Println("Write error:", err)
	}

}

func main() {
	listener, _ := net.Listen("tcp", ":7")

	defer listener.Close()

	for {

		conn, _ := listener.Accept()

		go handleConnection(conn)

	}

}
