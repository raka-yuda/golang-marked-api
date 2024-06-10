package main

import (
	"context"
	"fmt"
	"golang-marked-api/controllers"
	"golang-marked-api/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	// db     *dbCon.Queries
	ctx context.Context

	ProductController controllers.ProductController
	ProductRoutes     routes.ProductRoutes
)

func init() {
	ctx = context.TODO()

	ProductController = *controllers.NewProductController(ctx)
	ProductRoutes = routes.NewRouteProduct(ProductController)

	server = gin.Default()
}

func main() {
	// config, err := util.LoadConfig(".")

	// if err != nil {
	// 	log.Fatalf("failed to load config: %v", err)
	// }

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := server.Group("/api")

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "The contact API is working fine"})
	})

	ProductRoutes.ProductRoute(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	port := os.Getenv("PORT")

	log.Fatal(server.Run(":" + port))
}
