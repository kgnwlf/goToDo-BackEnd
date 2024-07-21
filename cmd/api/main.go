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

	fmt.Println(`

	________            ________            _______________________
	___  __/_____       ___  __ \_____      ___    |__  __ \___  _/
	__  /  _  __ \________  / / /  __ \     __  /| |_  /_/ /__  /
	_  /   / /_/ //_____/  /_/ // /_/ /     _  ___ |  ____/__/ /
	/_/    \____/       /_____/ \____/      /_/  |_/_/     /___/

	`)

	fmt.Println()

	fmt.Println("Starting Go To Do API...")

	err := http.ListenAndServe(":8000", r)

	if (err != nil) {
		log.Error(err)
	}

}
