package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"auth-system/database"
	"auth-system/handlers"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", handlers.RegisterHandler(db))
	mux.HandleFunc("POST /login", handlers.LoginHandler(db))

	mux.HandleFunc("GET /profile", handlers.JWTHandler(handlers.ProfilHandler()))
	mux.HandleFunc("/ws/matchmake", handlers.MatchmakingHandler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},

		AllowedMethods: []string{"GET", "POST", "OPTIONS"},

		AllowedHeaders: []string{"Authorization", "Content-Type"},

		Debug: false,
	})

	handler := c.Handler(mux)

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
