package services

import (
	"wardrobe-graphql/models"
	repository "wardrobe-graphql/repositories"

	"github.com/google/uuid"
)

type QuestionService struct {
	repo *repository.QuestionRepository
}

func NewQuestionService(repo *repository.QuestionRepository) *QuestionService {
	return &QuestionService{
		repo: repo,
	}
}

// Query
func (s *QuestionService) FindRandomAnsweredQuestion(limit int) ([]models.Question, error) {
	return s.repo.FindRandomAnsweredQuestion(limit)
}

// Mutation
func (s *QuestionService) CreateQuestion(question *models.Question) error {
	return s.repo.CreateQuestion(question)
}

func (s *QuestionService) DeleteQuestionById(id string) (bool, error) {
	// Validate Id
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return false, err
	}

	// Repo : Delete Question by id
	err = s.repo.DeleteQuestionById(uuidID)
	if err != nil {
		return false, err
	}

	return true, nil
}
