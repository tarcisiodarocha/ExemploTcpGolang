package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	ln, err := net.Listen("tcp", ":7788")
	checkError(err)
	fmt.Println("Server running!")
	for {
		conn, err := ln.Accept()
		checkError(err)
		go handleConn(conn)
	}

	os.Exit(0)
}

func handleConn(conn net.Conn) {
	fmt.Println("New Connection!")
	msg, err := bufio.NewReader(conn).ReadString('\n')
	checkError(err)
	conn.Write([]byte(msg))
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
