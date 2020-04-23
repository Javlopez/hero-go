package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	//SimpleRead()
	//ComplexReader()
	//ReadBooks()
	ReadBooksOption()

}

func ReadBooks() {
	url := "https://www.gutenberg.org/files/2701/old/moby10b.txt"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()

	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	//Create slice to hold the counts
	buckets := make([]int, 200)
	//Loop over the words
	for scanner.Scan() {
		n := HashBucketWord(scanner.Text()) //eg. Something... 65...123
		buckets[n]++
	}

	fmt.Println(buckets[65:122])
}

func ReadBooksOption() {
	url := "https://www.gutenberg.org/files/2701/old/moby10b.txt"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()

	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	//Create slice to hold the counts
	buckets := make([]int, 12)
	//Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12) //eg. Something... 65...123
		buckets[n]++
	}

	fmt.Println(buckets)
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

func HashBucketWord(word string) int {
	return int(word[0])
}

func ComplexReader() {
	url := "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}

		i++
	}

}

func SimpleRead() {
	url := "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bs))
}
