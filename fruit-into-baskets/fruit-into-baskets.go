package main

import "fmt"

func totalFruit(tree []int) int {
	max := 1
	m := make(map[int]int)
	i, j := 0, 0
	for j < len(tree) {
		if len(m) <= 2 {
			m[tree[j]] = j
			j++
		}
		if len(m) > 2 {
			min := len(tree) - 1
			for k := range m {
				if m[k] < min {
					min = m[k]
				}
			}
			i = min + 1
			delete(m, tree[min])
		}
		if j-i > max {
			max = j - i
		}
	}
	return max
}

func main() {
	tree := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(tree))
}
