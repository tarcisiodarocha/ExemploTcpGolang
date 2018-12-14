package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "www.ic.unicamp.br:80")
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("GET /~afalcao/ HTTP/1.0\r\n\r\n"))
	checkError(err)

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
