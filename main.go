package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yutaroyoshikawa/tipper-api/graph/generated"
	"github.com/yutaroyoshikawa/tipper-api/graph"
)

const defaultPort = ":8080"

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
          AllowOrigins: []string{os.Getenv("CORS_ALLOW_ORIGIN")},
	  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/health", func(c echo.Context) error {
	  return c.NoContent(http.StatusOK)
	})

	e.GET("/graphql", func(c echo.Context) error {
	  config := generated.Config{
	    Resolvers: &graph.Resolver{},
	  }
	  h := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	  h.ServeHTTP(c.Response(), c.Request())

	  return nil
	})

	e.GET("/", func(c echo.Context) error {
	  h := playground.Handler("GraphQL", "/graphql")

	  h.ServeHTTP(c.Response(), c.Request())

	  return nil
	})

	e.HideBanner = true

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}


	e.Logger.Fatal(e.Start(port))
}
