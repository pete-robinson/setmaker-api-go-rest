package services

import (
	"context"
	"errors"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/utils"
	m "setmaker-api-go-rest/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueSlug(t *testing.T) {
	tests := []struct {
		testName     string
		name         string
		initialSlugs []string
		expSlug      string
		count        int64
		expErr       error
		dbError      error
	}{
		{
			testName:     "Test a standard string",
			name:         "Some String",
			initialSlugs: []string{"some-string"},
			expSlug:      "some-string",
			count:        0,
		},
		{
			testName:     "Test a slug with non-standard characters",
			name:         "Sömê Śtrïñg",
			initialSlugs: []string{"some-string"},
			expSlug:      "some-string",
			count:        0,
		},
		{
			testName:     "Test a string with punctuation chars",
			name:         "Some/String",
			initialSlugs: []string{"some-string", ""},
			expSlug:      "some-string",
			count:        0,
		},
		{
			testName:     "Test a string with multiple punctuation chars",
			name:         "Some/-!+`String",
			initialSlugs: []string{"some-string", ""},
			expSlug:      "some-string",
			count:        0,
		},
		{
			testName:     "Test a slug which needs 1 iteration to complete",
			name:         "Some/String",
			initialSlugs: []string{"some-string", "some-string-1"},
			expSlug:      "some-string-1",
			count:        1,
		},
		{
			testName:     "Test a slug which needs 5 iteration to complete",
			name:         "Some/String",
			initialSlugs: []string{"some-string", "some-string-1", "some-string-2", "some-string-3", "some-string-4"},
			expSlug:      "some-string-4",
			count:        4,
		},
		{
			testName:     "Test a slug which fails on max-iterations",
			name:         "Some/String",
			initialSlugs: []string{"some-string", "some-string-1", "some-string-2", "some-string-3", "some-string-4", "some-string-5", "some-string-6", "some-string-7", "some-string-8", "some-string-9", "some-string-10"},
			expSlug:      "some-string-9",
			count:        10,
			expErr:       errors.New("Could not generate URL slug for Artist"),
		},
		{
			testName:     "Test a slug which fails on a DB error",
			name:         "Some/String",
			initialSlugs: []string{"some-string"},
			expSlug:      "some-string",
			count:        0,
			expErr:       errors.New("a DB error occurred"),
			dbError:      errors.New("a DB error occurred"),
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			ctx := context.TODO()

			a := &domain.Artist{
				Name: test.name,
			}

			artistRepo := &m.ArtistsRepository{}
			c := test.count
			for i := 0; i < len(test.initialSlugs); i++ {
				fs := utils.FieldSearch{
					Field: "slug",
					Query: test.initialSlugs[i],
				}

				artistRepo.On("Count", ctx, fs).Return(c, test.dbError)
				c--
			}

			service := NewArtistsService(artistRepo)

			e := service.uniqueSlug(ctx, a)
			assert.EqualValues(t, test.expSlug, a.Slug)

			if test.expErr != nil {
				assert.EqualValues(t, test.expErr, e)
			} else {
				assert.Nil(t, e)
			}
		})
	}
}
