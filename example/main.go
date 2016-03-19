package main

import (
	"net/http"
	"sync"

	"github.com/graphql-go/graphql"
	"github.com/mnmtanish/go-graphiql"
)

var (
	schema graphql.Schema
)

func init() {
	var (
		err error
		mtx sync.RWMutex
		msg = "Hello World"
	)

	RootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
		"getMessage": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				mtx.RLock()
				defer mtx.RUnlock()
				return msg, nil
			},
		},
	}}

	RootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: graphql.Fields{
		"setMessage": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"msg": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				mtx.Lock()
				defer mtx.Unlock()
				msg = p.Args["msg"].(string)
				return msg, nil
			},
		},
	}}

	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(RootQuery),
		Mutation: graphql.NewObject(RootMutation),
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.HandleFunc("/graphql", graphiql.ServeGraphQL(schema))
	http.ListenAndServe(":9001", nil)
}
