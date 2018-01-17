package main

import "fmt"

func changeMe(z []string){
	z[0] = "Pula"
	fmt.Println(z)
}

func main(){
	m := make([]string,2,25)
	fmt.Println(m)
	changeMe(m)

}
