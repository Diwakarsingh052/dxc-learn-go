package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"graphql-service/database"
	"graphql-service/graph"
	"log"
	"net/http"
)

const port = "8080"

func main() {

	db, err := database.ConnectToPostgres()
	if err != nil {
		log.Fatalln(err)
	}
	a, err := runAuth()
	if err != nil {
		log.Fatalln(err)
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db, Auth: a}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
