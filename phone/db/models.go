package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Phones struct {
	gorm.Model
	Name        string
	PhoneNumber string
}

func (p *Phones) GetNumber() string {
	return p.PhoneNumber
}

func (p *Phones) UpdateNumber(newNumber string) {
	p.PhoneNumber = newNumber
}

type ReadWriteNumber interface {
	GetNumber() string
	UpdateNumber(newNumber string)
}

func ConvertToReadWriteNumber(p []Phones) []ReadWriteNumber {
	var interfaceSlice []ReadWriteNumber

	for _, d := range p {
		interfaceSlice = append(interfaceSlice, &d)
	}

	return interfaceSlice
}

func Migrate() {
	// Migrate the schema
	// db.AutoMigrate(&Phones{})

	// Create
	// var phones []Phones
	// for _, number := range PhonesRaw {
	// 	var phoneForQuery Phones
	// 	phoneForQuery.PhoneNumber = number

	// 	phones = append(phones, phoneForQuery)
	// }

	// db.Create(&phones)

	// Read
	// Update - update product's price to 200
	//	db.Model(&product).Update("Price", 200)
	//	// Update - update multiple fields
	//	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//	// Delete - delete product
	//	db.Delete(&product, 1)
}

func GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=admin dbname=phones password=1"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetPhones(connection *gorm.DB) []Phones {
	var collectedPhones []Phones
	connection.Find(&collectedPhones)

	return collectedPhones
}
