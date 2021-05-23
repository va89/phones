package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDbPassword string = GetEnv("TEST_DB_PASSWORD", "1")

const (
	host   = "localhost"
	port   = 5432
	user   = "admin"
	dbname = "phones"
)

var DBConn = struct {
	Host   string
	Port   int
	User   string
	DBName string
}{
	host,
	port,
	user,
	dbname,
}

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

func UpdatePhonesField(connection *gorm.DB, field string, phone *Phones, value string) {

	connection.Debug().Model(phone).Select(field).Update(field, value)
	// fmt.Printf("%v#\n", res.Statement)
}

func BatchUpdatePhonesField(connection *gorm.DB, phones *[]Phones) {
	values := ""
	for i, phone := range *phones {
		values += fmt.Sprintf("(%d, '%s')", phone.ID, phone.PhoneNumber)
		if i < len(*phones)-1 {
			values += ","
		}
	}
	fmt.Println(values)
	connection.Exec(fmt.Sprintf(`update phones as phones set
					phone_number = phones2.phone_number
				    from (values
					%s
				    ) as phones2(id, phone_number)
				    where phones2.id = phones.id;`, values))
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

func DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", DBConn.Host, DBConn.Port, DBConn.User, DBConn.DBName, TestDbPassword)
}

func GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DSN()), &gorm.Config{})
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
