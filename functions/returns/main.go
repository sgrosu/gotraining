package main

import "fmt"

/*
func greet(fname,lname string) string {
	return fmt.Sprint(fname," ",lname)
}
*/

func greet(fname,lname string)(string, string) {
	return fname, lname
}
func main() {
	fmt.Println(greet("Ioan","Grosu"))
}