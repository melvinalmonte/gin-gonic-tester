package routes

import (
	"gin-gonic-tester/handlers"
	"github.com/gin-gonic/gin"
)

func Blueprint(v1 *gin.RouterGroup) {
	router := v1.Group("/v1")
	router.GET("/healthcheck", handlers.HealthCheck)
	router.GET("/posts", handlers.GetPosts)
}
