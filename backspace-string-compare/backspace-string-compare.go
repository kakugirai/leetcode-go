package main

import "fmt"

type stack []byte

func (s stack) Push(b byte) stack {
	return append(s, b)
}

func (s stack) Pop() stack {
	if len(s) == 0 {
		return s
	}
	return s[:len(s)-1]
}

func backspaceCompare(S string, T string) bool {
	var sstk stack
	for i := range S {
		if S[i] != '#' {
			sstk = sstk.Push(S[i])
		} else {
			sstk = sstk.Pop()
		}
	}
	var tstk stack
	for i := range T {
		if T[i] != '#' {
			tstk = tstk.Push(T[i])
		} else {
			tstk = tstk.Pop()
		}
	}

	if len(sstk) != len(tstk) {
		return false
	}
	for i := 0; i < len(sstk); i++ {
		if sstk[i] != tstk[i] {
			return false
		}
	}
	return true
}

func main() {
	//var s stack
	//var b byte
	//s = stack{1, 2, 3, 4}
	//s, b = s.Pop()
	//fmt.Println(s)
	//fmt.Println(b)
	S := "a#bc"
	T := "a#bc"
	fmt.Println(backspaceCompare(S, T))
}
