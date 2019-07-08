package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var curr []int
	//sort.Ints(candidates)
	dfs(candidates, 0, target, curr, &result)
	return result
}

func dfs(candidates []int, idx int, target int, curr []int, result *[][]int) {
	if target < 0 {
		return
	} else if target == 0 {
		cpy := make([]int, len(curr))
		copy(cpy, curr)
		*result = append(*result, cpy)
	} else {
		for i := idx; i < len(candidates); i++ {
			if len(curr) == 0 || candidates[i] >= curr[len(curr)-1] {
				curr = append(curr, candidates[i])
				dfs(candidates, idx, target-candidates[i], curr, result)
				curr = (curr)[:len(curr)-1]
			}
		}
	}
}

func main() {
	candidates := []int{2, 3, 5}
	fmt.Printf("%#v", combinationSum(candidates, 8))
}
