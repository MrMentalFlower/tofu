package server

import (
	"fmt"
	"net"
)

func Open() {
	conn, err := net.Listen("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println(conn, err)
		return
	}

	defer conn.Close()

	for {
		ackConn, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
		}

		var data net.Conn = ackConn

		fmt.Println(data)

	}

}

func listen() {

}
