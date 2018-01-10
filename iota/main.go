package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)
const (
	d = iota
	e
	f
)

const (
	_ = iota
	g = iota * 10
	h = iota * 10
)

const(
	_ = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println("binary\t\t\tdecimal")
	fmt.Printf("%b\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%T %v\n",0i,0i)
}

