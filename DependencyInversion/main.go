package main

import "fmt"

type (
	Database interface {
		Insert()
		Create()
		Update()
		Delete()
	}
	PostgresDB struct{}
	MockDB     struct{}
)

//PostgresDB
func NewPostgresDB() Database {
	return &PostgresDB{}
}

func (p *PostgresDB) Insert() {
	fmt.Println("PostgresDB Insert!")
}
func (p *PostgresDB) Create() {
	fmt.Println("PostgresDB Create!")
}
func (p *PostgresDB) Update() {
	fmt.Println("PostgresDB Update!")
}
func (p *PostgresDB) Delete() {
	fmt.Println("PostgresDB Delete!")
}

//MockDB
func NewMockDB() Database {
	return &MockDB{}
}
func (p *MockDB) Insert() {
	fmt.Println("MockDB Insert!")
}
func (p *MockDB) Create() {
	fmt.Println("MockDB Create!")
}
func (p *MockDB) Update() {
	fmt.Println("MockDB Update!")
}
func (p *MockDB) Delete() {
	fmt.Println("MockDB Delete!")
}

func InsertProduct(db Database) {
	db.Insert()
}
func CreateProduct(db Database) {
	db.Create()
}
func UpdateProduct(db Database) {
	db.Update()
}
func DeleteProduct(db Database) {
	db.Delete()
}

func main() {
	postgresDB := NewPostgresDB()
	DeleteProduct(postgresDB)

	mockDB := NewMockDB()
	InsertProduct(mockDB)

}
