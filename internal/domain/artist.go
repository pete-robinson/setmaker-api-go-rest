package domain

import (
	"strings"
	"unicode/utf8"

	"github.com/gosimple/slug"
	"github.com/pborman/uuid"
)

type Artist struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Slug  string    `json:"slug"`
	Image string    `json:"image"`
	Songs *[]Song   `json:"songs"`
}

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

func (a *Artist) CreateSlug(entropy string) string {
	strs := []string{a.Name, entropy}
	str := strings.Join(strs, " ")
	a.Slug = slug.Make(str)

	return a.Slug
}
