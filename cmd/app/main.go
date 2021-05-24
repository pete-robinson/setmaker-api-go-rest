package main

import (
	"net/http"
	"os"

	mdb "setmaker-api-go-rest/internal/database"
	"setmaker-api-go-rest/internal/handlers"
	"setmaker-api-go-rest/internal/repository"
	"setmaker-api-go-rest/internal/router"
	"setmaker-api-go-rest/internal/services"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	dbConn := buildConnection()
	db, err := mdb.ConfigureMongoClient(dbConn)
	if err != nil {
		panic(err)
	}

	// register repository
	artistsRepository := repository.NewArtistsRepository(db)

	// register services
	artistsService := services.NewArtistsService(artistsRepository)

	// register controllers
	controllers := map[string]router.RouteHandler{
		"artists": handlers.NewArtistsHandler(artistsService),
	}

	// build the router
	r := router.BuildRouter(controllers)

	// start the server
	log.Fatal(http.ListenAndServe(":8080", r))

}

func buildConnection() *mdb.ConnectionDto {
	conn := mdb.ConnectionDto{
		Url:  os.Getenv("MONGO_DB_URL"),
		User: os.Getenv("MONGO_DB_USER"),
		Pass: os.Getenv("MONGO_DB_PASS"),
		Name: os.Getenv("MONGO_DB_NAME"),
	}

	return &conn
}
