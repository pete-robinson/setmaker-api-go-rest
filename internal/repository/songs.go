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
	GetById(context.Context, uuid.UUID) (*domain.Song, error)
	Create(context.Context, *domain.Song) error
	Update(context.Context, *domain.Song) error
	Delete(context.Context, *domain.Song) error
	FindSongsByArtistId(context.Context, uuid.UUID) ([]*domain.Song, error)
}

type songsRepository struct {
	table string
	db    *mongo.Database
}

/**
 * create a new repository
 */
func NewSongsRepository(db *mongo.Database) *songsRepository {
	return &songsRepository{
		table: songsTable,
		db:    db,
	}
}

/**
 * find songs by artist id
 */
func (r *songsRepository) FindSongsByArtistId(ctx context.Context, id uuid.UUID) ([]*domain.Song, error) {
	cur, err := r.db.Collection(r.table).Find(ctx, bson.M{"artistId": id})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx) // defer close

	songs := make([]*domain.Song, 0)
	if err = cur.All(ctx, &songs); err != nil {
		return nil, err
	}

	return songs, nil
}

/**
 * Get a single song by ID
 */
func (r *songsRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Song, error) {
	var s *domain.Song

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

/**
 * Create a new song
 */
func (r *songsRepository) Create(ctx context.Context, s *domain.Song) error {
	s.ID = uuid.New()

	_, err := r.db.Collection(r.table).InsertOne(ctx, s)
	return err
}

/**
 * Update a song
 */
func (r *songsRepository) Update(ctx context.Context, s *domain.Song) error {
	update := bson.M{
		"$set": bson.M{
			"title":    s.Title,
			"artist":   s.Artist,
			"key":      s.Key,
			"tonality": s.Tonality,
		},
	}

	_, err := r.db.Collection(r.table).UpdateByID(ctx, s.ID, update)
	return err
}

/**
 * Delete a song
 */
func (r *songsRepository) Delete(ctx context.Context, s *domain.Song) error {
	_, err := r.db.Collection(r.table).DeleteOne(ctx, bson.M{"_id": s.ID})
	return err
}
