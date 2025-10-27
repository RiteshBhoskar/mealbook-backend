package main

import (
	"fmt"
	"net/http"

	"github.com/RiteshBhoskar/mealbook-backend/internal/db"
)

func main() {
	db.Connect()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Mealbook backend running!")
	})

	http.ListenAndServe(":8080", nil)
}
