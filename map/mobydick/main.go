package main

import (
	"fmt"
	"net/http"
	"log"
	"bufio"
)

func main () {
	// get the book Moby Dick

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		// scan the page
		scanner := bufio.NewScanner(res.Body)
		defer res.Body.Close()
		//set the split function for the scanning operation
		scanner.Split(bufio.ScanWords)
		// create slice to hold counts

		buckets := make([]int,200)
		//loop over the words
		for scanner.Scan() {
			n := HashBucket(scanner.Text(), 12)
			//fmt.Println(n)
			buckets[n]++
		}
		fmt.Println(buckets[:12])
		//fmt.Println(HashBucket("pula",12))
	}
}
func HashBucket(word string, nbuckets int) int{
	//letter := int(word[0])
	//bucket := letter % nbuckets
	//return bucket
	var sum int
	for _, letter := range word {
		sum+=int(letter)
	}
	return sum % nbuckets
}