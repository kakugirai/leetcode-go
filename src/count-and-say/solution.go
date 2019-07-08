package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	var builder strings.Builder
	counter := 0
	c := str[0]
	for i := range str {
		if str[i] == c {
			counter++
			continue
		}
		builder.WriteString(strconv.Itoa(counter))
		builder.WriteByte(c)
		c = str[i]
		counter = 1
	}

	builder.WriteString(strconv.Itoa(counter))
	builder.WriteByte(c)
	return builder.String()
}

func main() {
	n := 5
	fmt.Println(countAndSay(n))
}
