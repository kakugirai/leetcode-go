package main

import (
	"sort"
	"fmt"
)

func main() {
	strs := []string{"a", "b", "c"}
	sort.Strings(strs)
	fmt.Println("String", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}