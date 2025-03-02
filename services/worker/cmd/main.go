package main

import (
	"kafkacacheflow/services/worker/internal/handlers"
	"net/http"
)

func main() {
	hm := handlers.NewHandlerManager()

	http.HandleFunc("/change", hm.ChangeHandler)
	http.ListenAndServe(":8000", nil)
}
