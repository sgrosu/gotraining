package main

import "fmt"

func main() {
	var rez int
	for i:=0; i<1000; i++ {
		if i%3==0 || i%5==0 {
			rez += i
		}
	}
	fmt.Println(rez)
}
