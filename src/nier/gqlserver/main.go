package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rayhero/gqlgen-todos/src/base/gorm"
	"github.com/rayhero/gqlgen-todos/src/base/gqlgen/handlers"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = "localhost"
	port = "7777"
	gqlPath = "/graphql"
	gqlPgPath = "/"
	isPgEnabled = true
}

// Run spins up the server
func Run(orm *gorm.ORM) {
	// logger
	// log.Info("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))

	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	// Simple keep-alive/ping handler
	r.GET("/ping", Ping())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	// logger
	// log.Println("GraphQL @ " + endpoint + gqlPath)

	// Run the server
	// Inform the user where the server is listening
	log.Println("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatalln(r.Run(host + ":" + port))

}

func main() {
	gorm, err := gorm.Factory()
	if err != nil {
		log.Panic(err)
	}
	Run(gorm)
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
