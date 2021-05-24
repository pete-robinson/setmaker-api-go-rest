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

const Table = "artists"

// type ArtistsRepository interface {
// 	Get(id uuid.UUID) (*domain.Artist, error)
// 	Find(filter utils.QuerySort) ([]*domain.Artist, error)
// 	Create(*domain.Artist) error
// 	// Update(*domain.Artist) error
// 	// Delete(*domain.Artist) error
// }

type ArtistsRepository struct {
	table string
	db    *mongo.Database
}

func NewArtistsRepository(db *mongo.Database) *ArtistsRepository {
	return &ArtistsRepository{
		table: Table,
		db:    db,
	}
}

func (r *ArtistsRepository) Get(id uuid.UUID) (*domain.Artist, error) {
	var a *domain.Artist
	ctx := context.Background()

	found := r.db.Collection(r.table).FindOne(ctx, bson.M{"_id": id})
	if found == nil {
		log.Error(fmt.Sprintf("Artist %s not found", id))
		return nil, errors.New("Artist not found")
	}

	found.Decode(&a)

	return a, nil
}

func (r *ArtistsRepository) Find(filter *utils.QuerySort) ([]*domain.Artist, error) {
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

	defer res.Close(ctx) // close conn

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

func (r *ArtistsRepository) Count(queries ...utils.FieldSearch) (int64, error) {
	ctx := context.Background()
	filters := make([]bson.M, 0)

	for _, v := range queries {
		filters = append(filters, v.ToBson())
	}

	query := bson.M{"$and": filters}

	count, err := r.db.Collection(r.table).CountDocuments(ctx, query)
	return count, err
}

func (r *ArtistsRepository) Create(a *domain.Artist) error {
	ctx := context.Background()

	a.ID = uuid.New()
	_, err := r.db.Collection(r.table).InsertOne(ctx, a)
	return err
}

func (r *ArtistsRepository) Update(a *domain.Artist) error {
	ctx := context.Background()

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

func (r *ArtistsRepository) Delete(a *domain.Artist) error {
	ctx := context.Background()

	res, err := r.db.Collection(r.table).DeleteOne(ctx, bson.M{"_id": a.ID})
	fmt.Println(res)
	return err
}

// func (r *artistsRepository) Update(*domain.Artist) error {

// }

// func (r *artistsRepository) Delete(*domain.Artist) error {

// }
