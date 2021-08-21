package main

import "fmt"

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooChannel := make(chan int)

	go foo(fooChannel, 5)
	go foo(fooChannel, 2)

	v1 := <-fooChannel
	v2 := <-fooChannel
	fmt.Println(v1, v2)
}
