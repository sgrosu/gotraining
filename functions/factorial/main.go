package main

import "fmt"

func factorial(i int) int {
	if i == 0 {
		return 1
	} else {
		return i * factorial(i-1)
	}
}

func main(){
	fmt.Println(factorial(5))
}