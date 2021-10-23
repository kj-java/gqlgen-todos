package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gqlgen-todo/gqlgen-todos/graph"
	"github.com/gqlgen-todo/gqlgen-todos/graph/generated"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	// DBへの接続
	// host.docker.internalはコンテナをホストしてるipを指定してます
	// docker-composeでネットワークを作った方が良いとは思います。
	// なんか分からないけどうまく行かなかったのでとりあえずこうしてます。
	db, err := sql.Open("mypassword", "host=127.0.0.1 port=5432 user=user dbname=postgres sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

　　　　// Resolverに接続したdbを渡してResolverでdbにアクセスできるようにします。（Resolver.goに構造体を追加します）
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}