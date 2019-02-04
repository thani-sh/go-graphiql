package main

import (
	"log"
	"net/http"

	"github.com/alexsuslov/go-graphiql"
	"github.com/alexsuslov/go-graphiql/example/types"
	"github.com/graphql-go/graphql"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    types.RootQuery,
		Mutation: types.RootMutation,
	})
	if err != nil {
		panic(err)
	}

	go setMessage("Hello World")

	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.HandleFunc("/graphql", graphiql.ServeGraphQL(schema))
	http.ListenAndServe(":9001", nil)
}

func setMessage(msg string) {
	c, err := graphiql.NewClient("http://localhost:9001/graphql")
	if err != nil {
		panic(err)
	}

	q := `mutation _ { setMessage(msg: "` + msg + `") }`
	res, err := c.Send(&graphiql.Request{Query: q})
	if err != nil {
		panic(err)
	}

	if string(*res.Data) != `{"setMessage":"Hello World"}` {
		panic("bad response")
	}

	log.Println("listening on http://localhost:9001")
}

