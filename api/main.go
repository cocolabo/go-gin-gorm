package main

import (
	configs2 "github.com/cocolabo/go-gin-gorm/api/configs"
	database2 "github.com/cocolabo/go-gin-gorm/api/database"
	models2 "github.com/cocolabo/go-gin-gorm/api/models"
	repositories2 "github.com/cocolabo/go-gin-gorm/api/repositories"
	"log"
)

func main() {
	db, err := database2.Connect()

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models2.Contact{})

	contactRepository := repositories2.NewContactRepository(db)

	route := configs2.SetupRoutes(contactRepository)

	route.Run(":8000")

}
