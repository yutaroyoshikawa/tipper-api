package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yutaroyoshikawa/tipper-api/graph"
	"github.com/yutaroyoshikawa/tipper-api/graph/generated"
)

const defaultPort = "8080"
const authHeaderName = "Authorization"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Use(middleware.CORS())

	graphql := e.Group("/graphql")
	graphql.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:" + authHeaderName,
		Skipper: func(c echo.Context) bool {
			return c.Request().Header.Get(authHeaderName) == ""
		},
		Validator: func(idToken string, c echo.Context) (bool, error) {
			auth, err := app.Auth(c.Request().Context())
			if err != nil {
				return false, err
			}

			token, err := auth.VerifyIDToken(c.Request().Context(), idToken)

			if err != nil {
				return false, nil
			}

			c.Set("token", token)

			return true, nil
		},
	}))

	config := generated.Config{
		Resolvers: &graph.Resolver{},
	}
	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/", func(c echo.Context) error {
		h := playground.Handler("GraphQL", "/graphql")

		h.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	graphql.POST("", func(c echo.Context) error {
		c.Set("firebase", app)
		graphqlHandler.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	e.HideBanner = true

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
