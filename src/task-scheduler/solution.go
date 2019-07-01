package main

import (
	"fmt"
)

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	rec := make([]int, 26)
	for i := range tasks {
		rec[tasks[i]-'A']++
	}

	most := 0
	for i := range rec {
		if rec[i] > most {
			most = rec[i]
		}
	}

	idles := (most - 1) * (n + 1)
	for i := range rec {
		if most-1 < rec[i] {
			idles -= most - 1
		} else {
			idles -= rec[i]
		}
	}

	if idles > 0 {
		return len(tasks) + idles
	}
	return len(tasks)
}

func main() {
	task := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'D', 'D'}
	fmt.Println(leastInterval(task, 2))
}
