package repositories

import (
	"wardrobe-graphql/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DictionaryRepository struct {
	db *gorm.DB
}

func NewDictionaryRepository(db *gorm.DB) *DictionaryRepository {
	return &DictionaryRepository{
		db: db,
	}
}

// Query
func (r *DictionaryRepository) FindAllDictionaries() ([]models.Dictionary, error) {
	// Model
	var dictionaries []models.Dictionary

	// ORM
	err := r.db.Order("dictionary_type ASC").Find(&dictionaries).Error

	if err != nil {
		return nil, err
	}

	return dictionaries, nil
}

// Mutation
func (r *DictionaryRepository) CreateDictionary(dictionary *models.Dictionary) error {
	return r.db.Create(dictionary).Error
}

func (r *DictionaryRepository) DeleteDictionaryById(id uuid.UUID) error {
	return r.db.Delete(&models.Dictionary{}, "id = ?", id).Error
}
