package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered :", r)
	}
}

func say(s string) {
	defer wg.Done()
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
		if i == 2 {
			panic("Its a 222222!")
		}
	}
}

func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	say("there")
	wg.Wait()

	// Say bye after the wait
	wg.Add(1)
	say("Bye")
	wg.Wait()
}
