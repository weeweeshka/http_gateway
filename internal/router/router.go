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
		groupSso.POST("/login", handlers.Login())
	}

	groupTataisk := r.Group("/")
	{
		groupTataisk.POST("/create", handlers.CreateFilm())
		groupTataisk.GET("/:id", handlers.ReadFilm())
		groupTataisk.PATCH("/:id", handlers.UpdateFilm())
		groupTataisk.DELETE("/:id", handlers.DeleteFilm())
	}

	return r
}
