package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// postgres://xzfapyeo:Al1GhnCfoLOi43hHBtl8bKsf-03RgEFd@mouse.db.elephantsql.com/xzfapyeo
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to DB")
	} else {
		log.Printf("Connected to DB")
	}
}
