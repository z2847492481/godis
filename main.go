package main

import (
	"fmt"
	"net"
)

func main() {

	// listen

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("listening on port 6379")

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		resp := NewResp(conn)

		value, err := resp.Read()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// response client
		conn.Write([]byte("+OK\r\n"))
	}
}
