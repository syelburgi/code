package client_server

import (
	"bufio"
	"fmt"
	"net"
)

func Server() {
	listener, _ := net.Listen("tcp", ":8081")
	conn, _ := listener.Accept()
	for {
		str, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(str)
	}
}
