package main

import (
	"fmt"

	"github.com/Dejannnn/FiberCrudProj.git/database"
	"github.com/Dejannnn/FiberCrudProj.git/lead"
	"github.com/Dejannnn/FiberCrudProj.git/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Response struct {
	Message string `json:"message"`
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}
func main() {

	initDatabase()
	app := router.Router()
	app.Listen(":3000")
}
