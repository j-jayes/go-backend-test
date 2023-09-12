// @title Your API Name
// @version 1.0
// @description A brief description of your API.
// @host localhost:8080
// @BasePath /

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j-jayes/go-backend-test/pkg/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	/* port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default to localhost if not set
		log.Println("Defaulting to port 8080 as no PORT environment variable was set.")
	} */
	router := gin.Default()

	router.Static("/static", "../static")

	router.GET("/", func(c *gin.Context) {
		c.File("../static/index.html")
	})

	router.GET("/search/:query", api.SearchResults)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.Run("0.0.0.0:" + port)
	// run on port 8080
	router.Run("localhost:8080")
}
