package repositories

import (
	models2 "github.com/cocolabo/go-gin-gorm/api/models"
	"gorm.io/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Save(contact *models2.Contact) RepositoryResult {
	err := r.db.Save(contact).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: contact}
}

func (r *ContactRepository) FindAll() RepositoryResult {
	var contacts models2.Contacts

	err := r.db.Find(&contacts).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &contacts}
}

func (r *ContactRepository) FindOneById(id string) RepositoryResult {
	var contact models2.Contact

	err := r.db.First(&contact, "id = ?", id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &contact}
}
