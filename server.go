package graphiql

//go:generate go run html/gen.go

import (
	Data "github.com/alexsuslov/go-graphiql/html"
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
