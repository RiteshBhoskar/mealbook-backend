package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Mealbook backend running!")
	})

	http.ListenAndServe(":8080", nil)
}
