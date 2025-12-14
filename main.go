package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	database "github.com/rishabhrawatt/rssagg/neon/db"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello world")
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DBurl is not found")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("DB cannot connect", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	router := chi.NewRouter()
	//cors

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//routes

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", HandlerReadiness)
	v1Router.Get("/error", HandlerError)
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/getusers", apiCfg.middlewareAuth(apiCfg.HandlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString}

	fmt.Println("server starting on PORT:", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
