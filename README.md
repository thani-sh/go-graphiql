# go-graphiql

Go http handler to serve the graphiql in-browser IDE.

```go
package main

import (
	"net/http"

	"github.com/mnmtanish/go-graphiql"
)

func main() {
  schema := GetGraphQLSchema()
  http.HandleFunc("/graphql", graphiql.ServeGraphQL(schema))
	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.ListenAndServe(":9001", nil)
}
```
