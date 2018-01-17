package main

import "fmt"

func main() {
	var name string
	fmt.Println("Please enter your name at the prompt")
	fmt.Scan(&name)
	fmt.Println("Hello",name)
}
