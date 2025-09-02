package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weeweeshka/http_gateway/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	groupSso := r.Group("/auth")
	{
		groupSso.POST("/register", handlers.Register())
		groupSso.POST("/regapp", handlers.Regapp())
		groupSso.POST("/login", handlers.Login())
	}

	groupTataisk := r.Group("/films")
	{
		groupTataisk.POST("/create", handlers.CreateFilm())
		groupTataisk.GET("/read/:id", handlers.ReadFilm())
		groupTataisk.PATCH("/update/:id", handlers.UpdateFilm())
		groupTataisk.DELETE("/delete/:id", handlers.DeleteFilm())
	}

	groupComments := r.Group("/comments")
	{
		groupComments.POST("/create", handlers.CreateComment())
		groupComments.GET("/read/:id", handlers.ReadComment())
		groupComments.DELETE("/delete/:id", handlers.DeleteComment())
	}

	return r
}
