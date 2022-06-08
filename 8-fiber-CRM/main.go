package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/wralith/go-freecodecamp/8-fiber-CRM/database"
	"github.com/wralith/go-freecodecamp/8-fiber-CRM/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	initDatabase()
	setRoutes(app)
	app.Listen(8080)

	sqlDB, err := database.DBConn.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("lead.db"))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Successfully connected to DB")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("DB migrated")
}

func setRoutes(app *fiber.App) {
	app.Get("api/lead/:id", lead.GetLeadByID)
	app.Get("api/lead", lead.GetLeads)
	app.Post("api/lead", lead.NewLead)
	app.Delete("api/lead/:id", lead.DeleteLead)
}
