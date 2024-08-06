package main

import (
	"fmt"
	"net/http"
	"context"

	"goToDo-backEnd/internal/handlers"

	"github.com/go-chi/chi"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDoWorker struct {
	client *mongo.Client
}


func main () {
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println()
	fmt.Println("Starting Go To-Do API...")
	fmt.Println()

	fmt.Println(`

		________            ________            _______________________
		___  __/_____       ___  __ \_____      ___    |__  __ \___  _/
		__  /  _  __ \________  / / /  __ \     __  /| |_  /_/ /__  /
		_  /   / /_/ //_____/  /_/ // /_/ /     _  ___ |  ____/__/ /
		/_/    \____/       /_____/ \____/      /_/  |_/_/     /___/

	`)

	err := http.ListenAndServe(":8000", r)

	if (err != nil) {
		panic(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	fmt.Println("MongoDB has been connected.")
	fmt.Println()
	fmt.Println(client)

}
