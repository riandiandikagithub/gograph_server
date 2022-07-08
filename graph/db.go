package graph

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenConnection gorm
// reference : https://gorm.io/docs/connecting_to_the_database.html
func OpenConnection(host, port, username, password, dbName, timezone string) (db *gorm.DB, err error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable TimeZone=%s",
		host, port, username, password, dbName, timezone)

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return
	}

	return
}
