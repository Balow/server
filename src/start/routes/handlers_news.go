package routes

import (
	"encoding/json"
	"net/http"
	"start/models"
	"time"

	"github.com/gorilla/mux"
)

// NewsIndex : todo
func NewsIndex(w http.ResponseWriter, r *http.Request) {

	newss := []models.News{
		models.News{Id: "1", Title: "Hello1"},
		models.News{Id: "2", Title: "Hello2", Category: "a", Content: "a", Creation: time.Time{}},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(newss); err != nil {
		panic(err)
	}
}

// NewsShow : todo
func NewsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsID := vars["newsID"]

	news := models.News{Id: newsID, Title: "Hello1", Category: "a", Content: "a", Creation: time.Time{}}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(news); err != nil {
		panic(err)
	}

}
