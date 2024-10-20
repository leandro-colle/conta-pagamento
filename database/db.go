package database

import (
	"conta-pagamento/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	connectionData := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionData))
	if err != nil {
		log.Panic("Error connection")
	}
	DB.AutoMigrate(&models.Account{})
}
