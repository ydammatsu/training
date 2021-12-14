package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	x := 3
	y := 4
	z := add(x, y)
	fmt.Println(z)
}
