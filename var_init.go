package main

import (
	"fmt"
	"reflect"
)

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(reflect.TypeOf(c).String())
	fmt.Println(i, j, c, python, java)
}
