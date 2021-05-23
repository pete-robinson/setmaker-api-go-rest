package services

import (
	"context"
	"errors"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/utils"
	"strconv"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/uuid"
)

const Table = "artists" // DB table

type ArtistService struct {
	db    *mongo.Database
	table string
}

type ValidatorResponse map[string][]string

func NewArtistsService(db *mongo.Database) *ArtistService {
	return &ArtistService{
		db:    db,
		table: Table,
	}
}

func (svc *ArtistService) GetArtists(filter *utils.QuerySort) ([]*domain.Artist, error) {
	artists := make([]*domain.Artist, 0)
	ctx := context.TODO()

	// filter options
	opts := options.Find()
	opts.SetSort(bson.D{{Key: filter.Field, Value: filter.Operator}})

	// fetch results
	res, err := svc.db.Collection(svc.table).Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Error(err)
		return artists, err
	}

	defer res.Close(ctx) // close conn

	// loop cursor
	for res.Next(ctx) {
		var a domain.Artist
		err := res.Decode(&a)
		if err != nil {
			log.Error(err)
			return artists, err
		}

		artists = append(artists, &a)
	}

	return artists, nil
}

func (svc *ArtistService) CreateArtist(artist *domain.Artist) error {
	// set new id
	id := uuid.New()

	// ID created = set artist id
	artist.ID = id

	// create unique slug
	err := svc.uniqueSlug(artist)
	if err != nil {
		log.Error("Error generating URL slug")
		return err
	}

	// validate
	if errStr := artist.Validate(); len(errStr) > 0 {
		log.Error("Validation error", errStr)
		return errors.New(errStr[0]) // hacky but it'll do until i build a more robust abstraction for error mgmt
	}

	_, err = svc.db.Collection(svc.table).InsertOne(context.TODO(), artist)
	if err != nil {
		log.Error("Mongo error:", err)
		return err
	}

	return nil
}

func (svc *ArtistService) GetArtist(id uuid.UUID) (*domain.Artist, error) {
	var a *domain.Artist

	err := svc.db.Collection(svc.table).FindOne(context.TODO(), bson.M{"_id": id}).Decode(&a)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return a, nil

}

func (svc *ArtistService) uniqueSlug(a *domain.Artist) error {
	// loop through up to 10 times to create a unique slug
	var s string
	for i := 0; i < 20; i++ {
		if i == 0 {
			s = a.CreateSlug("")
		} else {
			ent := strconv.Itoa(i)
			s = a.CreateSlug(ent)
		}

		// call DB to see if slug is unique
		count, err := svc.db.Collection(svc.table).CountDocuments(context.TODO(), bson.M{"slug": s})
		if err != nil {
			return err
		}

		// slug is unique, we're done here
		if count == 0 {
			return nil
		}
	}

	return errors.New("Could not generate URL slug for Artist")
}
