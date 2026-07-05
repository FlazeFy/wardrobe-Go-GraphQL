package graph

import "wardrobe-graphql/services"

type Resolver struct {
	DictionaryService *services.DictionaryService
	FeedbackService   *services.FeedbackService
}
