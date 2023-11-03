package main

import (
	"log"
	"net/http"

	"github.com/k-rial/rest-api-golang/internal/taskstore"
)

type taskSever struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskSever {
	store := taskstore.New()
	return &taskSever{store: store}
}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
