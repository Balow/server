package routes

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"youk_back/entities"
	"youk_back/models"

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
	var outputData entities.News

	vars := mux.Vars(r)
	newsID := vars["newsID"]

	outputData = entities.GetNewsByID(newsID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputData); err != nil {
		panic(err)
	}

}

// NewsCreate : todo
func NewsCreate(w http.ResponseWriter, r *http.Request) {
	var inputData entities.News
	var outputData entities.News
	// var outputData models.News

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputData); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// var mappedData = entities.News{
	// 	Title:    inputData.Title,
	// 	Category: inputData.Category,
	// 	Content:  inputData.Content,
	// }

	outputData = entities.AddNews(inputData)
	// fmt.Println(outputData)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(outputData); err != nil {
		panic(err)
	}
}
