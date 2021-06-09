package main

import (
	"github.com/cocolabo/go-gin-gorm/configs"
	"github.com/cocolabo/go-gin-gorm/repositories"
	"log"

	"github.com/cocolabo/go-gin-gorm/database"
	"github.com/cocolabo/go-gin-gorm/models"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Contact{})

	contactRepository := repositories.NewContactRepository(db)

	route := configs.SetupRoutes(contactRepository)

	route.Run(":8000")

}
