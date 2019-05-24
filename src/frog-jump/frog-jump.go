package main

import "fmt"

func canCross(stones []int) bool {
	// check if the distance between two stones is more than k+1
	for i := 3; i < len(stones); i++ {
		if stones[i] > stones[i-1]*2 {
			return false
		}
	}

	// save stone positions to a hash set
	stonePositions := make(map[int]struct{})
	for i := range stones {
		stonePositions[stones[i]] = struct{}{}
	}

	lastStone := stones[len(stones)-1]
	var positions []int
	var jumps []int
	positions = append(positions, 0)
	jumps = append(jumps, 0)

	for len(positions) > 0 {
		// pop current position
		position := positions[len(positions)-1]
		positions = positions[:len(positions)-1]
		// pop jump distance
		jumpDistance := jumps[len(jumps)-1]
		jumps = jumps[:len(jumps)-1]

		for i := jumpDistance - 1; i <= jumpDistance+1; i++ {
			if i <= 0 {
				continue
			}
			nextPosition := position + i
			if nextPosition == lastStone {
				return true
			} else if _, ok := stonePositions[nextPosition]; ok {
				positions = append(positions, nextPosition)
				jumps = append(jumps, i)
			}
		}
	}
	return false
}

func main() {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	fmt.Println(canCross(stones))
}
