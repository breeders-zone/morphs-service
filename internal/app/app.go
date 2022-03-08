package app

import (
	"fmt"
	"log"

	"github.com/breeders-zone/morphs-service/internal/config"
	"github.com/breeders-zone/morphs-service/internal/handlers/http"
	"github.com/breeders-zone/morphs-service/internal/repositories"
	"github.com/breeders-zone/morphs-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// @title Morphs service API
// @version 1.0
// @description Morphs service API

// @host localhost:3000
// @BasePath /

// Run initializes whole application.
func Run() {
	app := fiber.New()

	conf, err := config.Init(".")
	if err != nil {
		log.Fatalf("not load config")
	}
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", conf.DBConfig.Server, conf.DBConfig.User, conf.DBConfig.Password, conf.DBConfig.Name)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	repos := repositories.NewRepositories(db)
	services := services.NewServices(repos)

	h := http.NewHandler(services)

	h.Init(app)

	app.Listen(":3000")
}