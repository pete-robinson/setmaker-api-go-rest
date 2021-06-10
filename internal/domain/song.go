package domain

import (
	"unicode/utf8"

	"github.com/google/uuid"
)

type Tonality int

const (
	TonalityMajor Tonality = iota
	TonalityMinor
	TonalityMixed
)

type Song struct {
	ID       uuid.UUID `bson:"_id,omitempty" json:"id"`
	Title    string    `bson:"name" json:"title"`
	Artist   uuid.UUID `bson:"artistId" json:"artistId"`
	Key      string    `bson:"key" json:"key"`
	Tonality Tonality  `bson:"tonality" json:"tonality"`
}

func (s *Song) Validate() []string {
	var errors []string

	if utf8.RuneCountInString(s.Title) == 0 {
		errors = append(errors, "Song title is required")
	}

	if s.Tonality > TonalityMixed || s.Tonality < TonalityMajor {
		errors = append(errors, "Invalid tonality")
	}

	return errors
}
