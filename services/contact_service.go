package services

import (
	"github.com/cocolabo/go-gin-gorm/dtos"
	"github.com/cocolabo/go-gin-gorm/models"
	"github.com/cocolabo/go-gin-gorm/repositories"
	"github.com/google/uuid"
	"log"
)

func CreateContact(contact *models.Contact, repository repositories.ContactRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	contact.ID = uuidResult.String()

	operationResult := repository.Save(contact)

	if operationResult.Result != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models.Contact)

	return dtos.Response{Success: true, Data: data}
}

func FindAllContacts(repository repositories.ContactRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models.Contacts)

	return dtos.Response{Success: true, Data: data}
}

func FindOneContactById(id string, repository repositories.ContactRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models.Contact)

	return dtos.Response{Success: true, Data: data}
}
