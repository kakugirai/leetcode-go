package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	arr := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]struct{})
	for _, word := range words {
		var curr string
		for i := range word {
			curr += arr[word[i]-'a']
		}
		m[curr] = struct{}{}
	}
	return len(m)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
