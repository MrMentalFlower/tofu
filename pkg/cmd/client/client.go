package client

import (
	"fmt"
	"net"
)

var message string

func Open() {
	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	fmt.Print("Input:")
	fmt.Scan(&message)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(conn, message+"\n")

	conn.Close()
}
