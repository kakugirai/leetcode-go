package main

import "fmt"

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[len(s)-i-1], s[i] = s[i], s[len(s)-i-1]
	}
}

func main() {
	s := []byte{'a', 'b', 'd', 'e'}
	reverseString(s)
	fmt.Println(s)
}
