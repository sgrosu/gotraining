package main

import (
	"fmt"
)

func main() {

	a := 43
	c := 1.9
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	var d *float64 = &c
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(*d)

	*d = 0.8
	fmt.Println(c)
}