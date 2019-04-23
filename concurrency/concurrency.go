package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go boring("sheep")
	go boring("fish")
	fmt.Println("I am listening")
	time.Sleep(time.Second * 2)
	fmt.Println("You are boring. I am leaving")
}
