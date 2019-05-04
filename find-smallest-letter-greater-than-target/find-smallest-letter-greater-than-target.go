package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1
	result := letters[0]
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			result = letters[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

func main() {
	fmt.Printf("%q", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'g'))
}
