package handlers

import (
	"github.com/99designs/gql-gen/handler"
	"github.com/gin-gonic/gin"

	"github.com/coy102/go-gql/internal/gql"
	"github.com/coy102/go-gql/internal/gql/resolvers"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler() gin.HandlerFunc {
    // NewExecutableSchema and Config are in the generated.go file
    c := gql.Config{
        Resolvers: &resolvers.Resolver{},
    }

    h := handler.GraphQL(gql.NewExecutableSchema(c))

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
    h := handler.Playground("Go GraphQL Server", path)
    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}