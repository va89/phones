package main

import (
	"github.com/va89/phone_normalizer/phone/db"
)

func main() {
	conn := db.GetDatabaseConnection()
	phones := db.GetPhones(conn)
	// numbersInterfaces := db.ConvertToReadWriteNumber(phones)
	db.NormalizeNumbers(&phones)

	db.BatchUpdatePhonesField(conn, &phones)

}
