package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
	Data string
}

func (v Vertex2) Setup() {
	v.Data = "Hello"
}

func (v Vertex2) Abs() float64 {
	fmt.Println("data", v.Data)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex2{
		X:    3,
		Y:    4,
		Data: "nohello",
	}
	v.Setup()
	fmt.Println(v.Abs())
}
