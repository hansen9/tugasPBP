package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	controllers "github.com/tubes/controllers"
)

func main() {
	router := mux.NewRouter()

	// Admin Routes
	router.HandleFunc("/admin/get_member", controllers.GetMemberByEmail).Methods("GET")
	router.HandleFunc("/admin/suspend_member/{email}", controllers.SuspendMember).Methods("PUT")
	router.HandleFunc("/admin/insert_film", controllers.InsertFilm).Methods("POST")
	router.HandleFunc("/admin/update_film/{id}", controllers.UpdateFilmById).Methods("PUT")

	// Member Routes
	router.HandleFunc("/member/register", controllers.Register).Methods("POST")
	router.HandleFunc("/member/update/{email}", controllers.UpdateMember).Methods("PUT")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3800"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(router)

	http.Handle("/", handler)
	fmt.Println("Connected to port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
