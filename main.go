package main

import (
	"log"
	"net/http"

	"learning/db"
	"learning/handlers"
	"learning/utils"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	h := &handlers.UserHandler{DB: db.DB}

	v1 := http.NewServeMux()

	v1.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetUsers(w, r)
		case http.MethodPost:
			h.CreateUser(w, r)
		default:
			utils.JSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})

	http.Handle("/v1/", http.StripPrefix("/v1", v1))

	http.ListenAndServe(":8080", nil)
}
