package main

import (
	"log"
	"wardrobe-graphql/config"
	"wardrobe-graphql/graph"
	"wardrobe-graphql/graph/generated"
	"wardrobe-graphql/repositories"
	"wardrobe-graphql/services"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	db := config.ConnectDatabase()

	// Repository
	repo := repositories.NewDictionaryRepository(db)

	// Service
	service := services.NewDictionaryService(repo)

	// Resolver
	resolver := &graph.Resolver{
		DictionaryService: service,
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver,
			},
		),
	)

	r := gin.Default()

	r.POST("/query", gin.WrapH(srv))

	r.GET("/playground",
		gin.WrapH(playground.Handler("GraphQL", "/query")),
	)

	r.Run(":8080")
}
