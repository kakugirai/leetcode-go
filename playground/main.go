package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("main dome")
	go func() {
		defer fmt.Println("goroutine1 done")
	}()
	fmt.Println("done")

}
