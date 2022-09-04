package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jwcarman/gomicro/api/authors"
	"github.com/jwcarman/gomicro/cmd/microservice/config"
	"github.com/jwcarman/gomicro/internal/database"
	"log"
)

// ...

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Instantiates the database
	postgres, err := database.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Instantiates the author service
	queries := database.New(postgres.DB)
	authorService := authors.NewService(queries)

	// Register our service handlers to the router
	router := gin.Default()
	fmt.Printf("Registering handlers...\n")
	authorService.RegisterHandlers(router)

	// Start the server
	fmt.Printf("Starting Gin server...\n")
	router.Run()
}
