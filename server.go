package main

import (
	"context"
	"fmt"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yutaroyoshikawa/tipper-api/graph"
	"github.com/yutaroyoshikawa/tipper-api/graph/generated"
	"github.com/yutaroyoshikawa/tipper-api/infrastructure"
	tipperMiddleware "github.com/yutaroyoshikawa/tipper-api/middleware"
	"github.com/yutaroyoshikawa/tipper-api/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	firebaseApp := service.InitializeFirebase(context.Background())

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Use(middleware.CORS())

	graphql := e.Group("/graphql")

	graphql.Use(tipperMiddleware.KeyAuth(firebaseApp))

	graphql.POST("", func(c echo.Context) error {
		resolverConfig := &graph.Resolver{
			Database: infrastructure.NewDatabase(firebaseApp),
		}
		if c.Get("token") != nil {
			resolverConfig.LoginUser = c.Get("token").(*auth.Token)
		}
		config := generated.Config{
			Resolvers: resolverConfig,
		}
		graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(config))
		graphqlHandler.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	e.GET("/", func(c echo.Context) error {
		h := playground.Handler("GraphQL", "/graphql")

		h.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	e.HideBanner = true

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
