package util

import "regexp"

func IsEmailFormat(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(emailRegex)

	return regex.MatchString(email)
}

func IsAlphanumeric(input string) bool {
	alphanumericRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(alphanumericRegex)

	return regex.MatchString(alphanumericRegex)
}
