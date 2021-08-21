package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}

func main() {
	fooChannel := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooChannel, i)
	}
	wg.Wait()
	close(fooChannel)

	for item := range fooChannel {
		fmt.Println(item)
	}

}
