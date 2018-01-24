package main

import "fmt"

func main (){

	myGreet := make(map[string]string)

	myGreet["Ion"] = "Neata!"
	myGreet["Maria"] = "Marsh!"
	myGreet["Roxana"] = "1"

	fmt.Println(myGreet)


	elements := map[string]map[string]string{
		"H": map[string]string{
			"name": "Hidrogen",
			"state": "gas",
		},
		"Fe": map[string]string{
			"name": "Iron",
			"state": "solid",
		},
	}
	fmt.Println(elements)
	if _, ok := elements["Ge"]; ok {
		fmt.Println("ok")
		} else {
		fmt.Println("nok")
	}

	for key, value := range elements {
		//fmt.Println(key, " - ", value)
		if value["name"] == "Hidrogen" {
			fmt.Println("mere", key)
		}
	}
}
