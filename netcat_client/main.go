package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	networkType = "tcp"
	port        = "8000"
	host        = "localhost"
)

func main() {
	conn, err := net.Dial(networkType, fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
