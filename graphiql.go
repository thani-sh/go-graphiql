package graphiql

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Request struct can be used when decoding HTTP requests sent by the GraphiQL
// IDE.
type Request struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}

// ServeGraphiQL is a handler function for HTTP servers
func ServeGraphiQL(res http.ResponseWriter, req *http.Request) {
	res.Write(content)
}

// ServeGraphQL returns a http.HandlerFunc which will process requests
// GraphQL queries and mutations sent from the GraphiQL editor.
func ServeGraphQL(s graphql.Schema) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

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
