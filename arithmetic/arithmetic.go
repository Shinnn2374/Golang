package main

import "fmt"

func main() {
	a := 5
	b := 10
	c := 50

	x := 2

	fmt.Println(a + b + c)
	fmt.Println(c - a - b)
	fmt.Println(c - a*b)
	fmt.Println(c/b - a)

	var val int = a % x
	fmt.Println(val)

	var y int = 8
	y++
	fmt.Println(y)

	var z int = 8
	z--
	fmt.Println(z)

}
