package main

import "fmt"

func main() {

	var a, b float64
	fmt.Println("please enter a small number")
	fmt.Scan(&a)
	fmt.Println("please enter a larger number")
	fmt.Scan(&b)
	fmt.Println("the result is:",b/a)

}
