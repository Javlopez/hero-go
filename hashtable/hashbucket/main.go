package main

import "fmt"

func main() {
	n := HashBucket("Golang", 12)
	fmt.Println(n)
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

func basicRunes() {
	//letter := 'A'
	//letter := rune("A"[0])
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}

func PrintLetters() {
	for i := 65; i < 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}
