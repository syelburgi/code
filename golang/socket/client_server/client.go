package client_server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Dialer() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the string")
		str, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, str + "\n")
	}
}
