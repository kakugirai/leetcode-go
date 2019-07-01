package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	curr := make([]byte, n*2)
	doGenerateParenthesis(n, n, 0, curr, &result)
	return result
}

func doGenerateParenthesis(left int, right int, idx int, curr []byte, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, string(curr))
		return
	}

	if left > 0 {
		curr[idx] = '('
		doGenerateParenthesis(left-1, right, idx+1, curr, result)
	}

	if right > 0 && left < right {
		curr[idx] = ')'
		doGenerateParenthesis(left, right-1, idx+1, curr, result)
	}
}

func main() {
	fmt.Println(generateParenthesis(2))
}
