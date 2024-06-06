package main

import (
	"context"
	"fmt"
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

	// ContactController controllers.ContactController
	// ContactRoutes     routes.ContactRoutes
)

func init() {
	ctx = context.TODO()

	// ContactController = *controllers.NewContactController(db, ctx)
	// ContactRoutes = routes.NewRouteContact(ContactController)

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
		ctx.JSON(http.StatusOK, gin.H{"message": "The contact APi is working fine"})
	})

	// ContactRoutes.ContactRoute(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	port := os.Getenv("PORT")

	log.Fatal(server.Run(":" + port))
}
