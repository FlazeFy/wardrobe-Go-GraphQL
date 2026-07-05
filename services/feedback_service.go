package services

import (
	repository "wardrobe-graphql/repositories"

	"github.com/google/uuid"
)

type FeedbackService struct {
	repo *repository.FeedbackRepository
}

func NewFeedbackService(repo *repository.FeedbackRepository) *FeedbackService {
	return &FeedbackService{
		repo: repo,
	}
}

func (s *FeedbackService) DeleteFeedbackById(id string) (bool, error) {
	// Validate Id
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return false, err
	}

	// Repo : Delete Feedback by id
	err = s.repo.DeleteFeedbackById(uuidID)
	if err != nil {
		return false, err
	}

	return true, nil
}
