package main

import "fmt"

func main() {

	var x [256]byte
	//fmt.Println(len(x))
	//fmt.Println(x[11])
	for i:=0; i<256; i++ {
		x[i] = byte(i)
	}
	fmt.Println(x)
}
