package main

import (
	"fmt"
	"math"
)

// formula t = - o / a - b => time equals minus the offset over the speed
// of the first tortoise - the speed of the second tortoise

func sround(input float64) float64{
	if math.IsNaN(input) {
		return math.NaN()
	}
	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}
	_, decimal := math.Modf(input)
	var rounded float64
	if decimal >= 0.5 {
		rounded = math.Ceil(input)
	} else {
		rounded = math.Floor(input)
	}
	return rounded * sign
}

func Race(v1, v2, g int) [3]int {
	s1 := float64(v1) / 60 / 60
	s2 := float64(v2) / 60 / 60

	t := -float64(g) / (s1-s2)
	fmt.Println(t)
	tr := sround(t)
	fmt.Println(sround(t))
	if t / 3600 < 1 {
		return [3]int{0,int(tr)/60,int(tr)%60}
	} else {

		return [3]int{int(tr)/3600,(int(tr)%3600)/60,(int(tr)%3600)%60}
	}


}

func main () {
	fmt.Println(Race(720,850,70))
}
