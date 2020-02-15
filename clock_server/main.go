package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const (
	networkType = "tcp"
	port        = "8000"
	host        = "localhost"
)

func main() {

	portInput := flag.String("port", port, "set port")
	flag.Parse()

	startServer(*portInput)
}

func startServer(portNumber string) {
	listener, err := net.Listen(networkType, fmt.Sprintf("%s:%s", host, portNumber))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[Starting TCP server listen in port: %s]\n", portNumber)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	fmt.Printf("A new connection has been started from: %s ....\n", c.RemoteAddr())
	for {
		output := time.Now().Format("15:04:05\n")
		_, err := io.WriteString(c, output)
		if err != nil {
			fmt.Printf("Error the client has been disconnected: %s\n", err.Error())
			fmt.Printf("Waiting for a new connection.....\n")
			return //since err this should be disconnected
		}
		fmt.Printf(output)
		time.Sleep(1 * time.Second)
	}
}
