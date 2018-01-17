package main

import "fmt"

func visit(nums []int, callback func(int)) {
	for _, num := range nums {
		callback(num)
	}
}

func main(){
	visit([]int{1,2,3,4},func(n int) {
		fmt.Println(n)
	})
}