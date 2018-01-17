package main

import "fmt"

func iseven(i float64) (float64, bool) {
	return i/2, int(i)%2 == 0
}

func main(){
	fmt.Println(iseven(7))
}