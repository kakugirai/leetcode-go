package main

import "fmt"

func generate(numRows int) [][]int {
	var triangle [][]int

	if numRows == 0 {
		return triangle
	}
	triangle = append(triangle, []int{1})

	for rowNum := 1; rowNum < numRows; rowNum++ {
		var row []int
		prevRow := triangle[rowNum-1]
		row = append(row, 1)
		for j := 1; j < rowNum; j++ {
			row = append(row, prevRow[j-1]+prevRow[j])
		}
		row = append(row, 1)
		triangle = append(triangle, row)
	}
	return triangle
}

func main() {
	fmt.Println(generate(4))
}
