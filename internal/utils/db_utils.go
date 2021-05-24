package utils

import "go.mongodb.org/mongo-driver/bson"

type QuerySort struct {
	Field    string
	Operator int
}

type FieldSearch struct {
	Field string
	Query interface{}
}

func (f *FieldSearch) ToBson() bson.M {
	return bson.M{f.Field: f.Query}
}
