package main

import "fmt"

func main() {

	for i:=50; i<170; i++ {
		//fmt.Println(i," - ", string(i), " - ", []byte(string(i)))
		fmt.Printf("%v %v %v\n", i, string(i), []byte(string(i)))
	}
}
