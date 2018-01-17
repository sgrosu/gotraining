package main

import "fmt"

func maxsl(alist ...int) int {
	maxx := alist[0]
	for i:=1; i < len(alist); i++ {
		if alist[i] > maxx {
			maxx = alist[i]
		}
	}
	return maxx
}

func main(){
	fmt.Println(maxsl(1,4,7,2,8))
}
