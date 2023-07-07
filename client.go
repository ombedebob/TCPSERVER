package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:55555")
	if err != nil {
		fmt.Println("Error dialing server", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Specify name first? ")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("Type message, (type Q to quit): ")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedClient + " says " + trimmedInput))
		if err != nil {
			fmt.Println("Error writing message", err.Error())
		}
	}
}
//bobombede