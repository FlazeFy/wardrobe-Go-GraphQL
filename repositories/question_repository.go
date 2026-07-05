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

func (r *QuestionRepository) CreateQuestion(question *models.Question) error {
	// Default
	question.ID = uuid.New()
	question.IsShow = false
	question.Answer = nil

	// Query
	return r.db.Create(question).Error
}
