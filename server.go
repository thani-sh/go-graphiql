package graphiql

//go:generate go run html/gen.go

import (
	"encoding/json"
	Data "github.com/alexsuslov/go-graphiql/html"
	"github.com/graphql-go/graphql"
	"github.com/mnmtanish/go-graphiql"
	"html/template"
	"net/http"
)

var (
	Tpl = template.New("index")
	Endpoint = "/graphql"
	T, _ = Tpl.Parse(Data.Content)
	)


// ServeGraphiQL is a handler function for HTTP servers
func ServeGraphiQL(res http.ResponseWriter, req *http.Request) {
	T.Execute(res, map[string]string{"Endpoint":Endpoint})
}

// ServeGraphQL create  handler function
func ServeGraphQL(s graphql.Schema) http.HandlerFunc {
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