package handlers

import (
	"encoding/json"
	"net/http"

	// "goToDo-backEnd/internal/tools"
	log "github.com/sirupsen/logrus"
	// "github.com/gorilla/schema"

)

// func getAllToDos (w http.ResponseWriter, r *http.Request) {

// 	var decoder *schema.Decoder = schema.NewDecoder()
// 	var err error

// 	err = decoder.Decode(&params, r.URL.Query())

// 	if err != nil {
// 		log.Error(err)
// 	}

// }

func sayHello (w http.ResponseWriter, r *http.Request) {
	var err error
	var response string = "Henlo friend."

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
	}

}
