package main

import (
	"graphql-service/database"
	"graphql-service/graph"
	"graphql-service/mid"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
)

const port = "8080"

func main() {

	db, err := database.ConnectToPostgres()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("main : Started : Initializing authentication support")
	a, err := runAuth()
	if err != nil {
		log.Fatalln(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Db: db, A: a}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//adding FetchToken as middleware
	//it means each time /query endpoint is called the FetchToken will run before it to fetch the token from the auth header
	http.Handle("/query", mid.FetchToken(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
