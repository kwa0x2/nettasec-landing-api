package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	db, err := gorm.Open(postgres.Open("postgres://nettasec:n1t2t3a4s5e6c@localhost:5435/nettasec-landing?sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	start := time.Now()
	for sqlDB.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connection database after 10 seconds")
			break
		}
	}

	fmt.Println("Connected to database: ", sqlDB.Ping()  == nil)
	DB = db
}
