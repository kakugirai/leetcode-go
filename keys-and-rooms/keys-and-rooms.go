package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	visited[0] = true

	var stk []int
	stk = append(stk, 0)
	for len(stk) > 0 {
		keys := rooms[stk[len(stk)-1]]
		stk = stk[:len(stk)-1]
		for i := range keys {
			if _, ok := visited[keys[i]]; !ok {
				visited[keys[i]] = true
				stk = append(stk, keys[i])
			}
		}
	}
	return len(visited) == len(rooms)
}

func main() {
	//rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	rooms := [][]int{{1}, {2}, {3}, {}}
	fmt.Println(canVisitAllRooms(rooms))
}
