package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dach3r/vecker/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manager is the function for administrate routers*/
func Manager() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.VerifyDB(routers.Register).Methods("POST"))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
