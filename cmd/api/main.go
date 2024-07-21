package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"goToDo-backEnd/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main () {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go To Do API...")

	err := http.ListenAndServe(":8000", r)

	if (err != nil) {
		log.Error(err)
	}

}
