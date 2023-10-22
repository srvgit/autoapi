package main

import (
	"autoapi/graph"
	"autoapi/service"
	"autoapi/store"
	"autoapi/util"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlHandler(resolver *graph.Resolver) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	boltStorer := store.NewBoltStore("configurations.db") // You should provide the path to your BoltDB file

	serverConfigService := service.NewServerConfigService(boltStorer)

	resolver := &graph.Resolver{
		ServerConfigService: serverConfigService,
	}

	r := gin.Default()
	r.POST("/query", graphqlHandler(resolver))
	r.GET("/", playgroundHandler())
	r.GET("/healthz", util.LivenessProbe)
	r.GET("/readiness", util.ReadinessProbe)
	r.Run()
}
