package main

import "fmt"

func main() {

	for i := 0; i<10; i++ {
		for j:= 0; j<10; j++ {
			fmt.Println("j - i = ",j-i)
			if i == 5 {
				fmt.Println("am ajuns la 5")
				break
			}
		}
	}
}