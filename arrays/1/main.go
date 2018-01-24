package main

import (
	"fmt"
	"strconv"

)

func main() {
	var anArray [10]string
	fmt.Println(anArray)
	for i:=0; i<len(anArray);i++{
		anArray[i] = strconv.Itoa(i)
	}
	fmt.Println(anArray)
}
