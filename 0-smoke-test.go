package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	buff := make([]byte, 1024)

	for {

		n, e := conn.Read(buff)

		if e != nil {

			conn.Write(buff[:n])

			break
		}

		conn.Write(buff[:n])
	}

}

func SmokeTest(network string, address string) {
	listener, _ := net.Listen(network, address)

	defer listener.Close()

	fmt.Println("SmokeTest Server Started")

	for {

		conn, _ := listener.Accept()

		go handleConnection(conn)

	}

}
