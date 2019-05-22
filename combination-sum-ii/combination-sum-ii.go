package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	fmt.Println(candidates)
	var result [][]int
	var current []int
	findCombination(candidates, 0, target, current, &result)

	return result
}

func findCombination(candidates []int, index int, target int, current []int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	if target < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		if i == index || candidates[i] != candidates[i-1] {
			current = append(current, candidates[i])
			findCombination(candidates, i+1, target-candidates[i], current, result)
			current = current[:len(current)-1]
		}
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
