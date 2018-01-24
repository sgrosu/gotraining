package main

import "fmt"

func main() {
	aSlice := []int {1,2,3}
	anotherSlice := []int {4,5,6}

	fmt.Println(append(aSlice,anotherSlice...))
}
