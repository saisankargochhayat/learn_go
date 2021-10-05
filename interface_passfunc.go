/**
Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
**/
package main

import "fmt"

type RetryInterface interface {
	RetryLogic(f func(a, b string), a, b string, c interface{})
}

type T struct {
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) RetryLogic(f func(string, string, interface{}), a, b string, c interface{}) {
	fmt.Println("Hello")
	x, ok1 := c.([]string)
	y, ok2 := c.([]int)
	if ok1 == false {
		fmt.Println("Calling foo")
		f(a, b, x)
	} else if ok2 == false {
		fmt.Println("bar")
		f(a, b, y)
	}
	// f(a, b, newC[])

}

func foo(a, b string, c []string) {
	fmt.Println(a, b, c)
	// fmt.Println(c)
}

func bar(a, b string, d []int) {
	fmt.Println(a, b, d)
}

func main() {
	var i RetryInterface = T{}
	v := []string{"asdas", "asdsa"}
	w := []int{2, 3}
	
	i.RetryLogic(foo, "alpha", "beta", v)
	i.RetryLogic(bar, "alpha", "beta", w)
}
