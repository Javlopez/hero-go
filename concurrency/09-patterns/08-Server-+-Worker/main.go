package main

import (
	"bytes"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	ch := make(chan string)
	go pool(ch, 4)
	go server(l, ch)
	time.Sleep(10 * time.Second)
}

func server(l net.Listener, ch chan string) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, ch)
	}
}

func pool(ch chan string, n int) {
	wch := make(chan int)
	results := make(chan int)
	for i := 0; i < n; i++ {
		go logger(wch, results)
	}
	go parse(results)
	for {
		addr := <-ch
		l := len(addr)
		wch <- l
	}
}

func logger(wch chan int, results chan int) {
	for {
		data := <-wch
		data++
		results <- data
	}
}

func parse(results chan int) {
	for {
		fmt.Println(<-results)
	}
}

func handler(c net.Conn, ch chan string) {
	addr := c.RemoteAddr().String()
	ch <- addr
	time.Sleep(100 * time.Millisecond)
	result := fmt.Sprintf("[GOROUTINE: %d, addr: %s]", getGID(), addr)
	c.Write([]byte(result))
	c.Close()
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
