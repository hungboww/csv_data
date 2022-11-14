package postgres

import (
	"ads/config"
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

func ConnectDatabase(c *config.Config) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Postgres.PostgresqlHost, c.Postgres.PostgresqlPort, c.Postgres.PostgresqlUser, c.Postgres.PostgresqlPassword, c.Postgres.PostgresqlDbname)

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
