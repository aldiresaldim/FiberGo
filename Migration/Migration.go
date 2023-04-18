package Migration

import (
	"FiberGo/Database"
	"FiberGo/Model/Entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := Database.DB.AutoMigrate(&Entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
