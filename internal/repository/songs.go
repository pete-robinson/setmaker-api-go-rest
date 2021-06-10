package repository

import (
	"context"
	"errors"
	"fmt"
	"setmaker-api-go-rest/internal/domain"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const songsTable = "songs"

type SongsRepository interface {
	Get(id uuid.UUID) (*domain.Song, error)
	Create(s *domain.Song) error
}

type songsRepository struct {
	table string
	db    *mongo.Database
}

func NewSongsRepository(db *mongo.Database) *songsRepository {
	return &songsRepository{
		table: songsTable,
		db:    db,
	}
}

func (r *songsRepository) Get(id uuid.UUID) (*domain.Song, error) {
	var s *domain.Song
	ctx := context.Background()

	// find single Song
	found := r.db.Collection(r.table).FindOne(ctx, bson.M{"_id": id})
	if found == nil {
		msg := fmt.Sprintf("Song %q not found", id)
		log.Error(msg)
		return nil, errors.New(msg)
	}

	err := found.Decode(&s)
	if err != nil {
		log.Error(err)
		return s, errors.New(fmt.Sprintf("Error fetching song: %q", id))
	}

	return s, nil
}

func (r *songsRepository) Create(s *domain.Song) error {
	ctx := context.Background()

	s.ID = uuid.New()

	_, err := r.db.Collection(r.table).InsertOne(ctx, s)
	return err
}
