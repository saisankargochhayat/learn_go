// Multiple defer statements lead to last in first out.
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	defer fmt.Println("asdasd")

	defer fmt.Println("sai")

}
