package main

import "fmt"

func main() {
	a := 4
	var ptr *int
	ptr = &a
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", *ptr)
	fmt.Printf("%d\n", ptr)
}
