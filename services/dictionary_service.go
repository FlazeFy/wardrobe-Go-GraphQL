package services

import (
	"wardrobe-graphql/models"
	repository "wardrobe-graphql/repositories"
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
