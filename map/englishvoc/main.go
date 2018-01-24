package main

import (
	"fmt"
	"net/http"
	//"io/ioutil"
	"log"
	"bufio"
	"os"
)

func main () {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	} else {
		/* part1
		bs, _ := ioutil.ReadAll(res.Body)
		str := string(bs)
		fmt.Println(str)
		*/
		words := make(map[string]string)

		sc := bufio.NewScanner(res.Body)
		sc.Split(bufio.ScanWords)

		for sc.Scan() {
			words[sc.Text()] = ""
		}
		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input", err)
		}
		i:=0
		for k, _ := range words {
			fmt.Println(k)
			if i == 20 {
				break
			}
			i++
		}
		fmt.Println("length of dict: ",len(words))
	}
}
