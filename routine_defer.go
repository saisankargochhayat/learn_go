package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
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
