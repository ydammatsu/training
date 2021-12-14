package main

import (
	"fmt"
	"strconv"
)

func add_integer(x, y int) int {
	return x + y
}

func add_string(x, y string) string {
	return x + y
}

func main() {
	x := "1"
	y := "234"
	z := add_string(x, y)

	a, err := strconv.Atoi(z)
	b := 1111
	if err != nil {
		fmt.Println("error")
	} else {
		c := add_integer(a, b)
		fmt.Println(c)
	}
}
