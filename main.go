package main

import (
	"log"
	"net/http"
	"os"
	controllers "replika-golang-fiber/handlers/controllers"
	"replika-golang-fiber/handlers/credential"
	"replika-golang-fiber/initialize"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		// Hanya load .env jika tidak berjalan di Railway
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Load config
	config, err := initialize.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	// Initialize database
	initialize.ConnectDB(&config)

	// Port
	port := os.Getenv("PORT")

	// Initialize Chi router
	r := chi.NewRouter()

	// Middleware: CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           3000,
	}))

	// Routing API
	r.Route("/api", func(r chi.Router) {
		// Public routes
		r.Post("/register", credential.RegisterHandler)
		r.Post("/login", credential.LoginHandler)

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(credential.AuthMiddleware)

			r.Route("/notes", func(r chi.Router) {
				r.Post("/", controllers.PostNoteHandler)
				r.Post("/datas", controllers.GetNotesHandler)
				r.Put("/{id}", controllers.UpdateNoteHandler)
				r.Delete("/delete/{id}", controllers.DeleteNoteHandler)
				r.Get("/note/{id}", controllers.GetNoteByID)
				r.Post("/favorites", controllers.PostFavoriteHandler)
				r.Post("/favorites/get", controllers.GetFavoriteHandler)
			})
		})
	})

	// Start the server
	log.Printf("Server running on port %s", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
