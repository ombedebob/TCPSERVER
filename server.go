package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("[Starting server...]")
	listener, err := net.Listen("tcp", "127.0.0.1:55555")
	if err != nil {
		fmt.Println("Server listening error", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
		}

		go doStuff(conn)
	}

}

func doStuff(conn net.Conn) {
	buff := make([]byte, 512)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading data", err.Error())
		return
	}

	fmt.Printf("Received data: %v ", string(buff))

}
//@bobombede