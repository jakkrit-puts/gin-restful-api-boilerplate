package main

import (
	"os"

	"example.com/jakkrit/ginbackendapi/configs"
	"example.com/jakkrit/ginbackendapi/routes"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := setupRouter()
	router.Run("localhost:" + os.Getenv("APP_PORT"))
}

func setupRouter() *gin.Engine {
	// Load .env
	godotenv.Load(".env")

	// mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// connnect db
	configs.Connection()

	router := gin.Default()

	// limit request size
	var maxBytes int64 = 1024 * 1024 * 10 // 10MB
	router.Use(limits.RequestSizeLimiter(maxBytes))
	// Serving static files
	router.Static("/public/images/", "./public/images/")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	apiV1 := router.Group("/api/v1")
	routes.InitHomeRoutes(apiV1)
	routes.InitUserRoutes(apiV1)

	return router
}
