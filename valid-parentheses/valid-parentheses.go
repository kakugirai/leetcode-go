package main

import (
	"fmt"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() string {
	//l := len(s)
	return s[0]
}

func main() {
	v := "{}["
	//m := map[string]string{")": "(", "}": "{", "]": "["}
	s := make(stack, 0)
	s.Push("a")
	a := s.Pop()
	fmt.Println(a)
	fmt.Println(v[0])
	fmt.Println(isValid(v))
}

func isValid(str string) (bool, stack) {
	// create the stack to keep track of brackets
	s := make(stack, 0)
	// create a hash map to keep track of mappings
	m := map[string]string{")": "(", "}": "{", "]": "["}
	// for every brackets in the expressions
	for _, v := range str {
		// if the character is an closing bracket
		if _, ok := m[string(v)]; ok {
			var topElement string
			if len(s) != 0 {
				topElement = s.Pop()
			} else {
				topElement = "#"
			}
			if m[string(v)] != topElement {
				return false, s
			}
		} else {
			s.Push(string(v))
		}
	}
	if len(s) == 0 {
		return true, s
	} else {
		return false, s
	}
}
