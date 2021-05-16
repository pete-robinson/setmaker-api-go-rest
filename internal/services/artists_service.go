package services

import (
	"context"
	"errors"
	"setmaker-api-go-rest/internal/domain"
	"strconv"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pborman/uuid"
)

const Table = "artists"

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

func (svc *ArtistService) CreateArtist(artist *domain.Artist) (bool, error) {
	// set new id
	artist.ID = uuid.NewUUID()

	// create unique slug
	err := svc.uniqueSlug(artist)
	if err != nil {
		log.Error("Error generating URL slug")
		return false, err
	}

	// validate
	if err := artist.Validate(); len(err) > 0 {
		log.Error("Validation error", err)
		return false, errors.New(err[0]) // hacky but it'll do until i build a more robust abstraction for error mgmt
	}

	res, err := svc.db.Collection(svc.table).InsertOne(context.TODO(), artist)
	_ = res // dumping to keep the linting happy
	if err != nil {
		log.Error("Mongo error:", err)
		return false, err
	}

	return true, nil
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
