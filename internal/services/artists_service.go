package services

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppService struct {
	db *mongo.Database
}

func NewArtistsService(db *mongo.Database) *AppService {
	return &AppService{
		db: db,
	}
}

func (svc *AppService) CreateArtist() {
	fmt.Println("CREATE FUNC")
}
