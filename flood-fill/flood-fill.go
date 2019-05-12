package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(image [][]int, i int, j int, color int, newColor int) {
	if i < 0 || i >= len(image) || j < 0 || j >= len(image[i]) || image[i][j] != color {
		return
	}
	image[i][j] = newColor
	dfs(image, i+1, j, color, newColor)
	dfs(image, i-1, j, color, newColor)
	dfs(image, i, j+1, color, newColor)
	dfs(image, i, j-1, color, newColor)
}

func main() {
	image := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		//{0, 0, 0},
	}
	fmt.Println(floodFill(image, 0, 0, 2))
}
