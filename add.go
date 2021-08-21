package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// When both parameters share a type
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(43, 58))
}
