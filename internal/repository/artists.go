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
	Get(id uuid.UUID) (*domain.Artist, error)
	Find(filter *utils.QuerySort) ([]*domain.Artist, error)
	Create(*domain.Artist) error
	Count(...utils.FieldSearch) (int64, error)
	Update(*domain.Artist) error
	Delete(*domain.Artist) error
}

type artistsRepository struct {
	table string
	db    *mongo.Database
}

func NewArtistsRepository(db *mongo.Database) *artistsRepository {
	return &artistsRepository{
		table: artistsTable,
		db:    db,
	}
}

func (r *artistsRepository) Get(id uuid.UUID) (*domain.Artist, error) {
	var a *domain.Artist
	ctx := context.Background()

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

func (r *artistsRepository) Find(filter *utils.QuerySort) ([]*domain.Artist, error) {
	// init result slice
	artists := make([]*domain.Artist, 0)
	ctx := context.Background()

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

		artists = append(artists, a)
	}

	return artists, nil
}

func (r *artistsRepository) Count(queries ...utils.FieldSearch) (int64, error) {
	ctx := context.Background()
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

func (r *artistsRepository) Create(a *domain.Artist) error {
	ctx := context.Background()

	// spawn new UUID
	a.ID = uuid.New()
	// create
	_, err := r.db.Collection(r.table).InsertOne(ctx, a)
	return err
}

func (r *artistsRepository) Update(a *domain.Artist) error {
	ctx := context.Background()

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

func (r *artistsRepository) Delete(a *domain.Artist) error {
	ctx := context.Background()

	_, err := r.db.Collection(r.table).DeleteOne(ctx, bson.M{"_id": a.ID})
	return err
}
