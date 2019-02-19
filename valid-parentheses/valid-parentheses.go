package main

import "fmt"

func main() {
	fmt.Println(isValid("{[]}{}"))
}

func isValid(s string) bool {
	var strs []string
	m := map[string]string{")": "(", "}": "{", "]": "["}
	for _, v := range s {
		if _, ok := m[string(v)]; ok {
			if len(strs) != 0:
				topElement :=
		}
	}
	return false
}