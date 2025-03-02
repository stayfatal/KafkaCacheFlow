package main

import (
	"kafkacacheflow/services/gateway/internal/handlers"
	"net/http"
)

func main() {
	hm := handlers.NewHandlerManager()

	http.HandleFunc("/change", hm.ChangeHandler)
	http.ListenAndServe(":8080", nil)
}
