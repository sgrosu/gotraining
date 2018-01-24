package main

import "fmt"

func main () {

	records := make([][]string,0)
	student1 := make([]string,4)

	student1[0] = "Ioan"
	student1[1] = "Grosu"
	student1[2] = "66"
	student1[3] = "pescar"

	records = append(records,student1)

	fmt.Println(records)

}
