package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j := 0, len(people)-1

	result := 0
	for i <= j {
		if people[i]+people[j] > limit {
			j--
		} else {
			i++
			j--
		}
		result++
	}
	return result
}

func main() {
	people := []int{3, 5, 3, 4}
	limit := 5
	fmt.Println(numRescueBoats(people, limit))
}
