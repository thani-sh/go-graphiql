package graphiql

import (
	"io/ioutil"
	"net/http"
)

var Content []byte

// ServeGraphiQL is a handler function for HTTP servers
func ServeGraphiQL(res http.ResponseWriter, req *http.Request) {
	if len(Content) == 0 {
		var err  error
		Content, err = ioutil.ReadFile("index.html")
		if err != nil {
			panic(err)
		}
	}
	res.Write(Content)
}
