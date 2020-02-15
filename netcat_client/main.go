package main

import (
	"flag"
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

	portInput := flag.String("port", port, "set port")
	flag.Parse()
	startClient(*portInput)
}

func startClient(port string) {
	conn, err := net.Dial(networkType, fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[Listen TCP server in port: %s]\n", port)

	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
