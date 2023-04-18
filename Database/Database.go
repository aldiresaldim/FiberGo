package Database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	const POSTGRES = "host=localhost user=postgres password=aldizix6ZWY dbname=FiberGo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := POSTGRES
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to database")
}
