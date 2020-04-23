package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://www.gutenberg.org/files/2701/old/moby10b.txt"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", bs)
}
