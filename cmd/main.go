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
	auth := service.NewAuthService(
		account.NewMongoRepo(mongodb),
	)

	// init server
	params := handler.Params{
		Auth: auth,
	}
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
		DbUri:  "mongodb://dbxadmin2:Aopj0R89Zp3J@203.161.44.242:27017/",
		DbName: "gmetour",
	}

	if err := Run(cfg); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
}
