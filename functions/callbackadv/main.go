package main

import "fmt"

func filter(alist []int,callback func(int) bool) []int {
	xs := []int{}
	for _, v := range alist {
		if callback(v) {
			xs = append(xs, v)
		}
	}
	return xs
}

func main() {
	x := filter([]int{1,2,3,4},func(n int) bool{
		return n > 1
	})
	fmt.Println(x)
}


