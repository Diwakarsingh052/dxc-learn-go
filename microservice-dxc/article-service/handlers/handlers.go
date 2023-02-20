package handlers

import (
	"article-service/article"
	"encoding/json"
	"log"
	"net/http"
)

// Handler are used to inject dependency in handlers
type Handler struct {
	article.Service
}

// CreateArticle handler handles the request to add a new article
func (c *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {

	var na article.NewArticle

	//decode the json in the NewArticle struct
	err := Decode(r, &na)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//adding article in postgres and generating event
	a, err := c.AddArticle(na)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//responding user with success
	Respond(w, a, http.StatusOK)

}

func (c *Handler) GetArticles(w http.ResponseWriter, r *http.Request) {

	//creating data struct to store json request sent by another service
	var data struct {
		Email string `json:"email"`
	}

	//decoding data in the data struct
	err := Decode(r, &data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fetch all the articles for the user
	articles, err := c.ListArticles(data.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//returning the list of the articles
	Respond(w, articles, http.StatusOK)

}

// Decode converts json to struct by reading the request body
func Decode(r *http.Request, val any) error {

	//NewDecoder reads the body and convert the json and put the data in the struct
	err := json.NewDecoder(r.Body).Decode(&val)

	if err != nil {
		return err
	}

	return nil
}

// Respond converts the struct to json
func Respond(w http.ResponseWriter, data any, statusCode int) error {

	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	//setting http headers to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	//writing the json to the client
	w.Write(jsonData)

	return nil

}
