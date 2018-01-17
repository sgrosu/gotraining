package variadic

//import "fmt"

func Average(sf []float64) float64 {
	total := 0.
	for _, v := range sf {
		total += v
	}
	res := total/float64(len(sf))
	return res
}
/*
func main () {
	n := Average(1,5,32,7,4,76)
	fmt.Println(n)
}
*/