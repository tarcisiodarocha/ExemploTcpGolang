package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", ":7788")
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("TESTE\n"))
	checkError(err)
	// msg, err := bufio.NewReader(conn).ReadString('\n')
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		fmt.Print(string(buf[0:n]))
		if err != nil {
			break
		}
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
