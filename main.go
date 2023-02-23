package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// tc := &controller.TaskController{}
	router := mux.NewRouter()
	// router.HandleFunc("/", tc.GetTasks).Methods(http.MethodGet)
	log.Print(http.ListenAndServe("0.0.0.0:80", router))

}
