package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	row := []int{1}
	prevRow := getRow(rowIndex - 1)
	for i := 1; i < len(prevRow); i++ {
		row = append(row, prevRow[i-1]+prevRow[i])
	}
	row = append(row, 1)
	return row
}

func main() {
	fmt.Println(getRow(5))
}
