package main

import "fmt"

func foo(aSlice ...int){
	fmt.Println(aSlice)
}

func main () {
	aslice := []int{1,2,3,4}
	foo(aslice...)
	foo()
	foo(1,2)
}