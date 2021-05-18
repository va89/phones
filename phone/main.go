package main

import (
	"regexp"

	"github.com/va89/phone_normalizer/phone/db"
)

func main() {
	conn := db.GetDatabaseConnection()
	phones := db.GetPhones(conn)
	numbersInterfaces := db.ConvertToReadWriteNumber(phones)
	filterNonDigits(&numbersInterfaces)

	// Add inteface "Update phone to Phones struct to make it easier to use updatePhones and filterNonDigits"
}

// func getFilteredPhones(phones []db.Phones) {
// 	for _, phone := range(phones) {
// 		phone.PhoneNumber =
// 	}
// }

func NormalizeNumbers(numbers *[]db.Phones) {
	for _, number := range *numbers {
		re := regexp.MustCompile(`\D`)
		newNumber := re.ReplaceAllString(number.PhoneNumber, "")
		number.UpdateNumber(newNumber)
	}
}
