package main

import (
	"fmt"
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := make(map[string]struct{})
	for i := range banned {
		bannedMap[banned[i]] = struct{}{}
	}

	r, _ := regexp.Compile("[^a-zA-Z]")
	count := make(map[string]int)
	paragraphArr := strings.Fields(strings.ToLower(r.ReplaceAllString(paragraph, " ")))
	for i := range paragraphArr {
		if _, ok := bannedMap[paragraphArr[i]]; !ok {
			count[paragraphArr[i]]++
		}
	}

	result := ""
	for i := range count {
		if result == "" || count[i] > count[result] {
			result = i
		}
	}
	return result
}

func main() {
	paragraph := "abc abc? abcd the jeff!"
	banned := []string{"abc", "abcd", "jeff"}
	fmt.Println(mostCommonWord(paragraph, banned))
}
