package main

import "fmt"

func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

func main() {
	fmt.Println(mySqrt(9))
}
