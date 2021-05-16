package mongo

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionDto struct {
	Url, User, Pass, Name string
}

func (mc *ConnectionDto) GetDsn() string {
	builder := strings.Builder{}
	builder.WriteString("mongodb+srv://")
	builder.WriteString(mc.User)
	builder.WriteString(":")
	builder.WriteString(mc.Pass)
	builder.WriteString("@")
	builder.WriteString(mc.Url)
	builder.WriteString(mc.Name)
	builder.WriteString("?retryWrites=true&w=majority")

	result := builder.String()
	return result
}

func ConfigureMongoClient(conn *ConnectionDto) (*mongo.Database, error) {
	// create db context
	ctx, cancel := getDbContext(10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn.GetDsn()))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Info("Mongo connected")

	db := client.Database(conn.Name)
	if db == nil {
		return nil, fmt.Errorf("DB %s does not exist", conn.Name)
	}

	return db, nil
}

func getDbContext(t int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
}
