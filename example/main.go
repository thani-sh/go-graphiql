package main

import (
	"encoding/json"
	"log"
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

	go setMessage("Hello World")

	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.HandleFunc("/graphql", serveGraphQL(schema))
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

func serveGraphQL(s graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendError := func(err error) {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		req := &graphiql.Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			sendError(err)
			return
		}

		res := graphql.Do(graphql.Params{
			Schema:        s,
			RequestString: req.Query,
		})

		if err := json.NewEncoder(w).Encode(res); err != nil {
			sendError(err)
		}
	}
}
