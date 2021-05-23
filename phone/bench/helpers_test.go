package db_test

import (
	"testing"

	"github.com/va89/phone_normalizer/phone/db"
)

func BenchmarkSample(b *testing.B) {

	for n := 0; n < b.N; n++ {
		input := []db.Phones{
			{PhoneNumber: "1234567890"},
			{PhoneNumber: "123 456 7891"},
			{PhoneNumber: "(123) 456 7892"},
			{PhoneNumber: "(123) 456-7893"},
			{PhoneNumber: "123-456-7894"},
			{PhoneNumber: "(123)456-7892"},
		}

		db.NormalizeNumbers(&input)
	}
}

func BenchmarkSample2(b *testing.B) {

	for n := 0; n < b.N; n++ {
		input := []db.Phones{
			{PhoneNumber: "1234567890"},
			{PhoneNumber: "123 456 7891"},
			{PhoneNumber: "(123) 456 7892"},
			{PhoneNumber: "(123) 456-7893"},
			{PhoneNumber: "123-456-7894"},
			{PhoneNumber: "(123)456-7892"},
		}

		db.NormalizeNumbers2(&input)
	}
}
