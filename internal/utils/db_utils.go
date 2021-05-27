package utils

import "go.mongodb.org/mongo-driver/bson"

type QuerySort struct {
	Field    string
	Operator int
}

// used for a key:value search parameter on DB
type FieldSearch struct {
	Field string
	Query interface{}
}

// convert to BSON util
func (f *FieldSearch) ToBson() bson.M {
	return bson.M{f.Field: f.Query}
}
