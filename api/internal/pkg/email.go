package pkg

import (
	"regexp"
)

const (
	emailPattern = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`
)

var (
	emailRegex = regexp.MustCompile(emailPattern)
)

func match(text string, regex *regexp.Regexp) []string {
	p := regex.FindAllString(text, -1)
	return p
}

func ParseEmail(text string) []string {
	return match(text, emailRegex)
}

func ValidateEmailInput(emails []string) bool {
	emailStr := ""
	for _, email := range emails {
		emailStr += email + " "
	}
	return len(ParseEmail(emailStr)) == len(emails)
}
