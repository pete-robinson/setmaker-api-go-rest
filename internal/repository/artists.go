package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/utils"

	"github.com/google/uuid"
)

const artistsTable = "artists"

type ArtistsRepository interface {
	GetById(context.Context, uuid.UUID) (*domain.Artist, error)
	Find(context.Context, *utils.QuerySort) ([]*domain.Artist, error)
	Create(context.Context, *domain.Artist) error
	Count(context.Context, ...utils.FieldSearch) (int64, error)
	Update(context.Context, *domain.Artist) error
	Delete(context.Context, *domain.Artist) error
}

type artistsRepository struct {
	table string
	db    *mongo.Database
}

/**
 * create a new repository
 */
func NewArtistsRepository(db *mongo.Database) *artistsRepository {
	return &artistsRepository{
		table: artistsTable,
		db:    db,
	}
}

/**
 * Get a single artist by ID
 */
func (r *artistsRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Artist, error) {
	var a *domain.Artist

	// find single artist
	found := r.db.Collection(r.table).FindOne(ctx, bson.M{"_id": id})
	if found == nil {
		// no artist found
		msg := fmt.Sprintf("Artist %q not found", id)
		log.Error(msg)
		return nil, errors.New(msg)
	}

	// attempt to decude result into Artist struct
	err := found.Decode(&a)
	if err != nil {
		log.Error(err)
		return a, errors.New(fmt.Sprintf("Error fetching artist: %q", id))
	}

	return a, nil
}

/**
 * Find an artist based on a filter
 */
func (r *artistsRepository) Find(ctx context.Context, filter *utils.QuerySort) ([]*domain.Artist, error) {
	// init result slice
	artists := make([]*domain.Artist, 0)

	// filter options
	opts := options.Find()
	if filter != nil {
		opts.SetSort(bson.D{{Key: filter.Field, Value: filter.Operator}})
	}

	// fetch results
	res, err := r.db.Collection(r.table).Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Error(err)
		return artists, err
	}

	defer res.Close(ctx) // defer close conn

	// loop cursor
	for res.Next(ctx) {
		var a *domain.Artist
		err := res.Decode(&a)
		if err != nil {
			log.Error(err)
			return artists, err
		}
		// append result to artists slice
		artists = append(artists, a)
	}

	return artists, nil
}

/**
 * Fetch count of artists based on a filter
 */
func (r *artistsRepository) Count(ctx context.Context, queries ...utils.FieldSearch) (int64, error) {
	// init filters
	filters := make([]bson.M, 0)

	// build query from FieldSearch variadic options
	for _, v := range queries {
		filters = append(filters, v.ToBson())
	}

	// construct query
	query := bson.M{"$and": filters}

	count, err := r.db.Collection(r.table).CountDocuments(ctx, query)
	return count, err
}

/**
 * Create a new artist
 */
func (r *artistsRepository) Create(ctx context.Context, a *domain.Artist) error {
	_, err := r.db.Collection(r.table).InsertOne(ctx, a)
	return err
}

/**
 * Update an existing artist
 */
func (r *artistsRepository) Update(ctx context.Context, a *domain.Artist) error {
	// create bson update query
	update := bson.M{
		"$set": bson.M{
			"name":  a.Name,
			"slug":  a.Slug,
			"image": a.Image,
		},
	}

	_, err := r.db.Collection(r.table).UpdateByID(ctx, a.ID, update)
	return err
}

/**
 * Delete an existing artist
 */
func (r *artistsRepository) Delete(ctx context.Context, a *domain.Artist) error {
	_, err := r.db.Collection(r.table).DeleteOne(ctx, bson.M{"_id": a.ID})
	return err
}
