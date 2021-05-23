package db_test

import (
	"testing"

	"github.com/va89/phone_normalizer/phone/db"
)

func TestNormalizeNumbers(t *testing.T) {
	input := []db.Phones{
		{PhoneNumber: "1234567890"},
		{PhoneNumber: "123 456 7891"},
		{PhoneNumber: "(123) 456 7892"},
		{PhoneNumber: "(123) 456-7893"},
		{PhoneNumber: "123-456-7894"},
		{PhoneNumber: "(123)456-7892"},
	}

	expected := []db.Phones{
		{PhoneNumber: "1234567890"},
		{PhoneNumber: "1234567891"},
		{PhoneNumber: "1234567892"},
		{PhoneNumber: "1234567893"},
		{PhoneNumber: "1234567894"},
		{PhoneNumber: "1234567892"},
	}

	t.Run("Test NormalizeNumbers", func(t *testing.T) {
		db.NormalizeNumbers(&input)

		for i := range input {
			if expected[i].PhoneNumber != input[i].PhoneNumber {
				t.Errorf("expected: %s, actual: %s", expected[i].PhoneNumber, input[i].PhoneNumber)
			}
		}
	})
}
