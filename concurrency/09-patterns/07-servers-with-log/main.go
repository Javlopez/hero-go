package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":5001")
	if err != nil {
		panic(err)
	}

	ch := make(chan string)
	go server(l, ch)
	go logger(ch)
	time.Sleep(10 * time.Second)
}

func logger(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}

func server(l net.Listener, ch chan string) {
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("err")
			log.Fatal(err)
		}
		go handler(c, ch)
	}
}

func handler(c net.Conn, ch chan string) {
	ch <- c.RemoteAddr().String()
	c.Write([]byte("ok"))
	c.Close()
}
