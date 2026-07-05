package services

import (
	"wardrobe-graphql/models"
	repository "wardrobe-graphql/repositories"
)

type QuestionService struct {
	repo *repository.QuestionRepository
}

func NewQuestionService(repo *repository.QuestionRepository) *QuestionService {
	return &QuestionService{
		repo: repo,
	}
}

func (r *QuestionService) CreateQuestion(question *models.Question) error {
	return r.repo.CreateQuestion(question)
}
