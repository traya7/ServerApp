package main

import (
	"ServerApp/domain"
	"ServerApp/domain/account"
	"ServerApp/handler"
	"ServerApp/service"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Config struct {
	Addr string

	DbUri  string
	DbName string
}

func Run(cfg Config) error {
	// init db connection
	mongodb, err := domain.NewMongoDB(cfg.DbUri, cfg.DbName)
	if err != nil {
		return err
	}
	log.Println("- DB connected!")

	// init services
	params := handler.Params{
		Auth: service.NewAuthService(account.NewMongoRepo(mongodb)),
	}

	// init server
	mux := handler.NewRouter(params)

	// add middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	// start server
	log.Println("- Server started!")
	return http.ListenAndServe(cfg.Addr, c.Handler(mux))
}

func main() {
	cfg := Config{
		Addr:   ":8080",
		DbUri:  "mongodb://myUserAdmin:hxadmin567%3F%3F@66.29.142.144:27017/",
		DbName: "gmeapp",
	}

	if err := Run(cfg); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
}
