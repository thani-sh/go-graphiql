package graphiql_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/alexsuslov/go-graphiql"
)

func TestResolv(t *testing.T) {
	options := graphiql.ResolverOptions{
		Name: "user",
		Imports: []string{
			"gitlab.42do.ru/servicechain/monitor/internal/db",
			"gopkg.in/mgo.v2/bson",
			"time"},
		Fields: [][]string{
			{"ID", "bson.ObjectId", `bson:"_id,omitempty"`},
			{"Name", "*string", `json:"name"`},
			{"FName", "*string", `json:"fullname"`},
			{"Phone", "*string", `json:"phone" `},
			{"Email", "*string", `json:"email" `},

			{"GroupID", "bson.ObjectId", `json:"group_id" bson:"group_id"`},
		},
	}
	buf := bytes.Buffer{}

	err := graphiql.GenResolver(&buf, options)
	if err != nil {
		panic(err)
	}
	log.Println(buf.String())

}
