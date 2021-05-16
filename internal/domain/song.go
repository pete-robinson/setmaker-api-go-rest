package domain

import (
	"github.com/pborman/uuid"
)

type Tonality int

const (
	Tonality_Major Tonality = iota
	Tonality_Minor
	Tonality_Mixed
)

type Song struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Artist   *Artist   `json:"artist"`
	Key      string    `json:"key"`
	Tonality Tonality  `json:"tonality"`
}
