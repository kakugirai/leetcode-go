package main

import "fmt"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) interface{} {

	tmp := make([]int32, len(q))
	for i := range tmp {
		tmp[i] = int32(i + 1)
	}
	for i, v := range q {
		if v <= tmp[i] {
			continue
		}
		if v+1-int32(i) > 2 {
			return "Too chaotic"
		}
		// delete that element
		tmp = append(tmp[:i], tmp[i+1:]...)

	}
}

func main() {
	a := []int32{1, 2, 3, 4, 5, 6, 7}
	a = append(a[:2], append([]int32{a[6]}, a[2:]...)...)
	fmt.Println(a)
}
