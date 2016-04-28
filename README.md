# go-graphiql

Go http handler to serve the graphiql in-browser IDE.

```go
package main

import (
	"net/http"

	"github.com/mnmtanish/go-graphiql"
)

func main() {
	http.HandleFunc("/graphql", myGraphQLHandler)
	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.ListenAndServe(":9001", nil)
}
```
