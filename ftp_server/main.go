package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"
)

const (
	networkType = "tcp"
	host        = "localhost"
	portNumber  = 3000
)

func main() {
	fmt.Printf("main\n")

	startFTPServer()
}

func startFTPServer() {
	listener, err := net.Listen(networkType, fmt.Sprintf("%s:%d", host, portNumber))
	fmt.Printf("[Starting FTP server listen in port: %d]\n", portNumber)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	fmt.Printf("A new connection has been started from: %s ....\n", c.RemoteAddr())
	for {
		command, err := ioutil.ReadAll(c)
		if err != nil || string(command) == "" {
			//log.Print(err)
			continue
		}

		cmd, err := commandHandler(string(command))
		log.Print(string(command))
		if err != nil {
			continue
		}
		fmt.Printf("Command received: %+v\n", cmd)
		time.Sleep(10 * time.Second)
	}
}

func commandHandler(command string) (string, error) {
	switch command {
	case "ls":
		return "list of files", nil
	case "cd":
		return "up directory", nil
	}

	return "", fmt.Errorf("The command %s is not supportted", command)
}
