package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1 / x * myPow(1/x, -(n+1))
	} else if n == 2 {
		return x * x
	} else if n%2 == 0 {
		return myPow(myPow(x, n/2), 2)
	}
	return x * myPow(myPow(x, n/2), 2)
}

func main() {
	x := 2.0
	fmt.Println(myPow(x, 3))
}
