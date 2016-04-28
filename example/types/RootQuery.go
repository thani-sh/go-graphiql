package types

import (
	"github.com/graphql-go/graphql"
	"github.com/mnmtanish/go-graphiql/example/logic"
)

// RootQuery ...
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getMessage": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				msg := logic.GetMessage()
				return msg, nil
			},
		},
	},
})
