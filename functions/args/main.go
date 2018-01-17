package main

import "fmt"

func main (){
	greet("Ionel","Grosu")
	println(greet)
}

func greet(fname,lname string) {
	fmt.Println("Hello",fname,lname)
}
