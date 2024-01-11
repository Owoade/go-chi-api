package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Owoade/go-chi-api/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *db.Queries
}

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {

		log.Fatal("Port not provided")

	}

	db_url := os.Getenv("DB_URL")

	if db_url == "" {
		log.Fatal("DB_URL not provided")
	}

	conn, err := sql.Open("postgres", db_url)

	api_config := ApiConfig{
		DB: db.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", handler)

	router.Get("/error", error_handler)

	router.Post("/user", api_config.create_author_handler)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	log.Printf("Server starting on %v", port)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Couldn't start server")
	}

}
