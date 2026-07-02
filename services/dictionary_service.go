package services

import (
	"time"
	"wardrobe-graphql/models"
	repository "wardrobe-graphql/repositories"

	"github.com/google/uuid"
)

type DictionaryService struct {
	repo *repository.DictionaryRepository
}

func NewDictionaryService(repo *repository.DictionaryRepository) *DictionaryService {
	return &DictionaryService{
		repo: repo,
	}
}

func (s *DictionaryService) FindAllDictionaries() ([]models.Dictionary, error) {
	return s.repo.FindAllDictionaries()
}

func (s *DictionaryService) CreateDictionary(dictionaryType, dictionaryName string) (*models.Dictionary, error) {
	// Prepare model
	dictionary := &models.Dictionary{
		ID:             uuid.New(),
		DictionaryType: dictionaryType,
		DictionaryName: dictionaryName,
		CreatedAt:      time.Now(),
	}

	// Repo : Create dictionary
	err := s.repo.CreateDictionary(dictionary)

	if err != nil {
		return nil, err
	}

	return dictionary, nil
}
