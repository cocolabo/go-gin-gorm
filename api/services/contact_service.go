package services

import (
	dtos2 "github.com/cocolabo/go-gin-gorm/api/dtos"
	models2 "github.com/cocolabo/go-gin-gorm/api/models"
	repositories2 "github.com/cocolabo/go-gin-gorm/api/repositories"
	"github.com/google/uuid"
	"log"
)

func CreateContact(contact *models2.Contact, repository repositories2.ContactRepository) dtos2.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	contact.ID = uuidResult.String()

	operationResult := repository.Save(contact)

	if operationResult.Result != nil {
		return dtos2.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models2.Contact)

	return dtos2.Response{Success: true, Data: data}
}

func FindAllContacts(repository repositories2.ContactRepository) dtos2.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos2.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models2.Contacts)

	return dtos2.Response{Success: true, Data: data}
}

func FindOneContactById(id string, repository repositories2.ContactRepository) dtos2.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos2.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models2.Contact)

	return dtos2.Response{Success: true, Data: data}
}
