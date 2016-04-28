package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/mnmtanish/go-graphiql"
	"github.com/mnmtanish/go-graphiql/example/types"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    types.RootQuery,
		Mutation: types.RootMutation,
	})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.HandleFunc("/graphql", graphiql.ServeGraphQL(schema))
	http.ListenAndServe(":9001", nil)
}
