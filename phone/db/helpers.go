package db

import (
	"bytes"
	"os"
	"regexp"
)

func NormalizeNumbers(numbers *[]Phones) {
	for i := range *numbers {
		re := regexp.MustCompile(`\D`)
		filtered := re.ReplaceAllString((*numbers)[i].PhoneNumber, "")
		(*numbers)[i].PhoneNumber = filtered
	}
}

// 20 times faster that regex version
func NormalizeNumbers2(numbers *[]Phones) {
	for i := range *numbers {
		// re := regexp.MustCompile(`\D`)
		var buf bytes.Buffer

		for _, ch := range (*numbers)[i].PhoneNumber {
			if ch <= '0' && ch <= '9' {
				buf.WriteRune(ch)
			}
		}
		// filtered := re.ReplaceAllString((*numbers)[i].PhoneNumber, "")
		(*numbers)[i].PhoneNumber = buf.String()
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
