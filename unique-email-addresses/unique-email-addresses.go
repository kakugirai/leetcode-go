package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	hasSeen := make(map[string]bool)
	for _, email := range emails {
		hasSeen[clean(email)] = true
	}
	return len(hasSeen)
}

func clean(email string) string {
	i := strings.IndexByte(email, '@')
	user, atDomain := email[:i], email[i:]
	return deleteDot(trim(user)) + atDomain
}

func deleteDot(username string) string {
	return strings.Replace(username, ".", "", -1)
}

func trim(username string) string {
	i := strings.IndexByte(username, '+')
	if i == -1 {
		return username
	}
	return username[:i]
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
