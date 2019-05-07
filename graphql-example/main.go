package main

import (
	"github.com/graphql-go/graphql"
	gh "github.com/graphql-go/handler"
	"log"
	"net/http"
	"server_common/graphql-example/conf"
	"server_common/graphql-example/db"
	"server_common/graphql-example/handler"
)


func main(){

	confErr := conf.InitConfig("dev")
	if confErr != nil {
		log.Fatal(confErr)
	}

	db.Init()

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    handler.QueryType,
		Mutation: handler.MutationType,
	})

	schemaGraphiql, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    handler.QueryType,
		Mutation: handler.MutationType,
	})

	h := gh.New(&gh.Config{
		Schema:   &schemaGraphiql,
		Pretty:   true,
		GraphiQL: true,
	})
	log.Println("server start, port:", 1323)
	http.Handle("/graphql", handler.Handler(schema))
	http.Handle("/graphiql", h)
	log.Fatal(http.ListenAndServe(":1323", nil))
}