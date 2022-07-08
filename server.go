package main

import (
	"graph_server/graph"
	"graph_server/graph/generated"
	"log"
	"net/http"
	"os"

	"graph_server/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	Database, err := graph.OpenConnection(
		utils.GetEnv("WARPIN_HOST", "127.0.0.1"),
		utils.GetEnv("WARPIN_PORT", "5432"),
		utils.GetEnv("WARPIN_USERNAME", "postgres"),
		utils.GetEnv("WARPIN_PASSWORD", "postgres"),
		utils.GetEnv("WARPIN_DBNAME", "db_warpin"),
		utils.GetEnv("WARPIN_TIMEZONE", "Asia/Jakarta"),
	)

	if err != nil {
		log.Fatal("err connect db")
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: Database}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
