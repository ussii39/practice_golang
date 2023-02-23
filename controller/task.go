package controller

import (
	"database/sql"
	"fmt"
	"net/http"
)

type TaskController struct {
	Db *sql.DB
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "a")
}
