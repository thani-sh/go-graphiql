package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/mnmtanish/go-graphiql"
	"github.com/mnmtanish/go-graphiql/example/types"
)

// Request struct can be used to decode HTTP requests
// sent by the GraphiQL IDE.
type Request struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    types.RootQuery,
		Mutation: types.RootMutation,
	})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.HandleFunc("/graphql", serveGraphQL(schema))
	http.ListenAndServe(":9001", nil)
}

func serveGraphQL(s graphql.Schema) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		Err := func(errMsg string) {
			res.WriteHeader(500)
			res.Write([]byte(errMsg))
		}

		request := &Request{}
		if err := json.NewDecoder(req.Body).Decode(request); err != nil {
			Err("Internal Error")
			return
		}

		p := graphql.Params{Schema: s, RequestString: request.Query}
		o := graphql.Do(p)
		if o.HasErrors() {
			Err("Internal Error")
			return
		}

		result, err := json.Marshal(o)
		if err != nil {
			Err("Internal Error")
			return
		}

		res.Write(result)
	}
}
