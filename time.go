package main

import (
	"fmt"
	"time"
)

func run() {
	fmt.Println("hello")
	cooldown := 10 * .2
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("bye")
	fmt.Println(time.Duration(cooldown))
}

func main() {
	run()
}
