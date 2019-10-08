package main

import "fmt"

func reverseVowels(s string) string {
	m := map[byte]byte{
		'a': 'a',
		'i': 'i',
		'u': 'u',
		'e': 'e',
		'o': 'o',
		'A': 'A',
		'I': 'I',
		'U': 'U',
		'E': 'E',
		'O': 'O',
	}
	i, j := 0, len(s)-1
	for i < j {
		_, oki := m[s[i]]
		_, okj := m[s[j]]
		if !oki {
			i++
		}
		if !okj {
			j--
		}
		if oki && okj {
			temp := s[i]
			s = s[:i] + string(s[j]) + s[i+1:]
			s = s[:j] + string(temp) + s[j+1:]
			i++
			j--
		}
	}
	return s
}

func main() {
	s := "lEetcode"
	fmt.Println(reverseVowels(s))
}
