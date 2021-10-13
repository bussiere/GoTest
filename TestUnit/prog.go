package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func main() {
	var result int
	result = Sum(5, 5)
	fmt.Println(result)
}
