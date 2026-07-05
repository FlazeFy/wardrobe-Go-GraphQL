package repositories

import (
	"wardrobe-graphql/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepository {
	return &FeedbackRepository{
		db: db,
	}
}

func (r *FeedbackRepository) DeleteFeedbackById(id uuid.UUID) error {
	return r.db.Delete(&models.Feedback{}, "id = ?", id).Error
}
