package main

import (
	"graphql-crud/graph"
	"graphql-crud/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var todos []*model.Todo
	var users []*model.User

	users = append(users,
		&model.User{
			ID:   "001",
			Name: "John Doe",
		},
		&model.User{
			ID:   "002",
			Name: "Jane Doe",
		},
	)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		TodoList: todos,
		UserList: users,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
