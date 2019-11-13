package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"

	"github.com/rayhero/gqlgen-todos/src/base/gorm"
	"github.com/rayhero/gqlgen-todos/src/base/gqlgen"
	"github.com/rayhero/gqlgen-todos/src/base/gqlgen/resolvers"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(orm *gorm.ORM) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	c := gqlgen.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm,
		},
	}

	h := handler.GraphQL(gqlgen.NewExecutableSchema(c))

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
