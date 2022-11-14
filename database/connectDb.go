package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "172.24.0.2"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "postgres"
)

func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("ConnectDatabase success!")

	DB = db
}
