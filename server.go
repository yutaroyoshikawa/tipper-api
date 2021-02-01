package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yutaroyoshikawa/tipper-api/graph"
	"github.com/yutaroyoshikawa/tipper-api/graph/generated"
)

const defaultPort = ":80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Use(middleware.CORS())

	config := generated.Config{
	  Resolvers: &graph.Resolver{},
	}
	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	e.GET("/health", func(c echo.Context) error {
	  return c.NoContent(http.StatusOK)
	})

	e.GET("/", func(c echo.Context) error {
	  h := playground.Handler("GraphQL", "/graphql")

	  h.ServeHTTP(c.Response(),c.Request())

	  return nil
	})

	e.POST("/graphql", func(c echo.Context) error {
	  graphqlHandler.ServeHTTP(c.Response(), c.Request())

	  return nil
	})

	e.HideBanner = true

	e.Logger.Fatal(e.Start(port))
}
