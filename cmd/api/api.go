package api

import (
	"encoding/json"
	"net/http"
)

type AddToDoParams struct {
	toDo string
}

type ModifyToDoParams struct {
	toDoID int64 // To Do should exist and have an ID.
	modifyTo string // Change To Do to this.
}

type DeleteToDoParams struct {
	toDoID string
}

type ToDos struct {
	id int64
	toDo string
}

type ToDoResponse struct {
	// Response code. 201 when creating a toDo, acknowledge all other successful calls with 200.
	code int

	// Optional fields.
	message string // Optional success message.
	allToDos []ToDos // Optional list of toDos.
}

type Error struct {
	code int
	message string
}
