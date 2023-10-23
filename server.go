package main

import (
	"autoapi/config"
	"autoapi/graph"
	"autoapi/service"
	"autoapi/store"
	"autoapi/util"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	vehicleGraphletFeatures := os.Getenv("GRAPHLET_FEATURES")
	features := config.ParseFeatures(vehicleGraphletFeatures)

	boltStorer := store.NewBoltStore("configurations.db") // You should provide the path to your BoltDB file

	serverConfigService := service.NewServerConfigService(boltStorer)

	resolver := &graph.Resolver{
		ServerConfigService: serverConfigService,
	}

	if features["VEHICLE"] {
		sugar.Infow("Vehicle feature is enabled")
		resolver.VehicleService = (*service.VehicleService)(service.NewVehicleService(boltStorer))
	}
	if features["DEALER"] {
		sugar.Infow("Vehicle feature is enabled")
		resolver.DealerService = (*service.DealerService)(service.NewDealerService(boltStorer))
	}

	r := gin.Default()
	r.POST("/query", graphqlHandler(resolver))
	r.GET("/", playgroundHandler())
	r.GET("/healthz", util.LivenessProbe)
	r.GET("/readiness", util.ReadinessProbe)
	r.Run(":9090")
}
