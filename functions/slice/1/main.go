package main

import (

	"github.com/sgrosu/gotraining/functions/variadic"
	"fmt"
)

func sl(li []float64) {
	for i,v := range li {
		fmt.Println(i, "=", v)
	}
}

func main () {
	p := []float64{1:6,4,7}
	sl(p)
	fmt.Println(variadic.Average(p))
	fmt.Println(p)
}