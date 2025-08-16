package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weeweeshka/http_gateway/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	groupSso := r.Group("/")
	{
		groupSso.POST("/register", handlers.Register())
	}
}
