package pkg

import (
	"net/mail"
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

func ExtractEmail(text string) []string {
	return match(text, emailRegex)
}

func validMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

func ValidateEmailInput(emails []string) bool {
	for _, email := range emails {
		isValid := validMailAddress(email)
		if !isValid {
			return false
		}
	}
	return true
}
