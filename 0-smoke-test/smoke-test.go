package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {

	buff := make([]byte, 1024)
	for {

		_, e := conn.Read(buff)

		if e != nil {
			fmt.Println("CONN RESET")
			return
		}

		conn.Write(buff)

	}

}

func main() {
	listener, _ := net.Listen("tcp", ":8000")

	defer listener.Close()

	for {

		conn, _ := listener.Accept()

		go handleConnection(conn)

	}

}
