package main

import (
	"log"
	"rest-api-note-taking/config"
	"rest-api-note-taking/routes"

	"github.com/joho/godotenv"
	"github.com/swaggo/swag/example/override/docs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Blog API"
	docs.SwaggerInfo.Description = "This is Golang Backend Project Blogspot."
	docs.SwaggerInfo.Version = "1.0"

	// database connection setup
	db := config.ConnectDatabase()
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDb.Close()

	// route setup
	r := routes.SetupRouter(db)
	r.Run(":8080")
}
