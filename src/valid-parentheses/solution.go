package main

import "fmt"

type stack []byte

func (s stack) Push(v byte) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, byte) {
	return s[:len(s)-1], s[len(s)-1]
}

func isValid(str string) bool {
	// create the stack to keep track of brackets
	var s stack
	//s := make(stack, 0)
	// create a hash map to keep track of mappings
	m := map[byte]byte{')': '(', '}': '{', ']': '['}
	// for every brackets in the expressions
	for i := range str {
		// if the character is an closing bracket
		if _, ok := m[str[i]]; ok {
			// get the top byte in the stack
			var topByte byte
			if len(s) != 0 {
				s, topByte = s.Pop()
			} else {
				// placeholder
				topByte = '#'
			}
			// if the closing bracket is not our top byte
			if m[str[i]] != topByte {
				return false
			}
		} else {
			s = s.Push(str[i])
		}
	}
	return len(s) == 0
}

func main() {
	str := "{}["
	fmt.Println(isValid(str))
}
