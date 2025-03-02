package handlers

import (
	"io"
	"log"
	"net/http"
)

type HandlersManager struct {
}

func NewHandlerManager() *HandlersManager {
	return &HandlersManager{}
}

func (hm *HandlersManager) ChangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		resp, err := http.Post("http://worker:8000/change", "application/json", r.Body)
		if err != nil {
			http.Error(w, "Internal serve error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, string(body), http.StatusInternalServerError)
				log.Println(err)
				return
			}

			http.Error(w, string(body), resp.StatusCode)
		}
	}
}
