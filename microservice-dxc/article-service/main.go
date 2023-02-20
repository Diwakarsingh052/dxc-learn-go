package main

import (
	"article-service/article"
	"article-service/database"
	"article-service/handlers"
	"log"
	"net/http"
)

func main() {
	//mux := chi.NewRouter()

	db, err := database.ConnectToPostgres()
	if err != nil {
		log.Fatalln(err)
	}
	r, err := database.ConnectToRedis()
	if err != nil {
		log.Fatalln(err)
	}

	//setting  postgres and redis dependency in service struct, so article package can access the db
	s := article.Service{DB: db, Client: r}

	h := handlers.Handler{Service: s}
	http.HandleFunc("/article/view", h.GetArticles)
	http.HandleFunc("/article/add", h.CreateArticle)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
