package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"fmt"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch url: %v\n", err)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch url reading: %v\n", err)
		}

		fmt.Printf("%s", b)
	}
}
