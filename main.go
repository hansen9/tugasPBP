package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	controllers "github.com/latihan/controllers"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// Routes

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3800"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(router)

	http.Handle("/", handler)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
