package main

import (
	"log"
	"logger/cmd/api/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB     *mongo.Client
	Logger controllers.Logger
}

func NewApp(mongo *mongo.Client) *App {
	return &App{
		DB:     mongo,
		Logger: controllers.NewLoggerController(mongo),
	}
}

func (a *App) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(middleware.Heartbeat("/healthCheck"))
	return mux
}

func (a *App) serve() {
	//define server
	srv := &http.Server{
		Addr:    ":80",
		Handler: a.routes(),
	}

	//start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
