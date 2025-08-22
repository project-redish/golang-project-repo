package main

import (
	_ "golang-project-repo/docs"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	port = ":8080"
)

// @title My Go API
// @version 1.0
// @description This is a sample API for demonstration.
// @host localhost:8080
// @BasePath /
func main() {

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	http.HandleFunc("/greet", greetingHandler)

	log.Println("Starting Server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server", err)

	}
}

// @Summary Get a greeting
// @Description Returns a simple greeting message
// @Tags greetings
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /greet [get]
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
