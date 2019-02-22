package main

import "fmt"

type stack []byte

func (s stack) Push(v byte) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, byte) {
	return s[:len(s)-1], s[len(s)-1]
}

func main() {
	str := "{}[]"
	fmt.Println(isValid(str))
}

func isValid(str string) (bool, stack) {
	// create the stack to keep track of brackets
	s := make(stack, 0)
	// create a hash map to keep track of mappings
	m := map[byte]byte{')': '(', '}': '{', ']': '['}
	// for every brackets in the expressions
	for i := range str {
		// if the character is an closing bracket
		if _, ok := m[str[i]]; ok {
			var topElement byte
			if len(s) != 0 {
				s, topElement = s.Pop()
			} else {
				topElement = '#'
			}
			if m[str[i]] != topElement {
				return false, s
			}
		} else {
			s.Push(str[i])
		}
	}
	return len(s) == 0, s
}
