package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://qiita.com/tnoda_/items/b503a72eac82862d30c6

func main() {
	fmt.Print("> ")
	var input1 string
	fmt.Scan(&input1)
	fmt.Println(input1)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	scanner.Scan()
	input2 := scanner.Text()
	fmt.Println(input2)
}
