package main
/*
function to determine a number's largest prime factor
 */

import (
	"fmt"
	//"math"
	"time"
)


func lpf(n int) int {
	i:=2
	//x := &n
	for i*i < n {

		for n % i == 0 {
			n = n / i
			//*x = n
		}
		//fmt.Println(n)
		i+=1

	}

	return n
}

func main(){
	start := time.Now()
	fmt.Println(lpf(600851475143))
	fmt.Println(time.Since(start))
}
