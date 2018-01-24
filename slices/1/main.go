package main

import "fmt"

func main() {
	//x := []string{"a","b","vc"}
	//fmt.Println(x[1:])

	y := make([]int,0,5)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	for i:=0;i<80;i++ {
		y  = append(y,i)
		fmt.Println("Length",len(y),"capacity", cap(y),"Value",y[i])
	}


}
