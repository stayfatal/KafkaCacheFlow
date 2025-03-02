package handlers

import (
	"net/http"
)

type HandlersManager struct {
}

func NewHandlerManager() *HandlersManager {
	return &HandlersManager{}
}

func (hm *HandlersManager) ChangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Error(w, "test error", http.StatusConflict)
	}
}
