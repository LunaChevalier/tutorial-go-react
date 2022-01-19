package main

import (
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/LunaChevalier/tutorial-go-react/graph"
	"github.com/LunaChevalier/tutorial-go-react/graph/generated"
)

const dataSource = "root:root@tcp(mysql:3306)/react-go-app?charset=utf8&parseTime=True&loc=Local"
const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
    db, err := gorm.Open("mysql", dataSource)
    if err != nil {
        panic(err)
    }
    if db == nil {
        panic(err)
    }
    defer func() {
        if db != nil {
            if err := db.Close(); err != nil {
                panic(err)
            }
        }
    }()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
