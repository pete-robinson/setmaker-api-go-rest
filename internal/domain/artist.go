package domain

import (
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

// core struct
type Artist struct {
	ID    uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string    `bson:"name" json:"name"`
	Slug  string    `bson:"slug" json:"slug"`
	Image string    `bson:"image" json:"image"`
	Songs *[]Song   `bson:"songs" json:"songs"`
}

// simple struct validation
// @todo use a library to validate this better
func (a *Artist) Validate() []string {
	var errors []string
	a.Name = strings.TrimSpace(a.Name)

	if utf8.RuneCountInString(a.Name) == 0 {
		errors = append(errors, "Name is required")
	}

	if utf8.RuneCountInString(a.Slug) == 0 {
		errors = append(errors, "Slug is required")
	}

	return errors
}

// CreateSlug creates a url-safe version of the artist title.
// Entropy argument used for auto-append.
func (a *Artist) CreateSlug(entropy string) string {
	strs := []string{a.Name, entropy}
	str := strings.Join(strs, " ")
	a.Slug = slug.Make(str)

	return a.Slug
}
