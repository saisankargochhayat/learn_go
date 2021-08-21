package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)
	pixel := make([]uint8, dx)
	for i := range data {
		for j := range pixel {
			pixel[j] = uint8((i ^ j) * 13)
		}
		data[i] = pixel
	}
	return data
}

func main() {
	pic.Show(Pic)
}
