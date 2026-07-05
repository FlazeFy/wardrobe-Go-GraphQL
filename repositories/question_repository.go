package repositories

import (
	"wardrobe-graphql/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

// Query
func (r *QuestionRepository) FindRandomAnsweredQuestion(limit int) ([]models.Question, error) {
	// Model
	var questions []models.Question

	// ORM
	err := r.db.Order("RANDOM()").Where("answer is not null").Limit(limit).Find(&questions).Error

	if err != nil {
		return nil, err
	}

	return questions, nil
}

// Mutation
func (r *QuestionRepository) CreateQuestion(question *models.Question) error {
	// Default
	question.ID = uuid.New()
	question.IsShow = false
	question.Answer = nil

	// Query
	return r.db.Create(question).Error
}

func (r *QuestionRepository) DeleteQuestionById(id uuid.UUID) error {
	return r.db.Delete(&models.Question{}, "id = ?", id).Error
}
