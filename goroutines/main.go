package main

import (
	"fmt"
	"log"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//f("direct")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
