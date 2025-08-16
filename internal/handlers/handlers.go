package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/weeweeshka/http_gateway/internal/domain/models"
	clients "github.com/weeweeshka/http_gateway/internal/grpcClients"
	pbSSO "github.com/weeweeshka/sso_proto/gen/go/sso"
	pbTataisk "github.com/weeweeshka/tataisk_proto/gen/go/tataisk"
	"net/http"
)

var (
	ssoClient, _     = clients.SetupGateway()
	_, tataiskClient = clients.SetupGateway()
)

func Register(email string, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		req := models.RegRequest{
			Email:    email,
			Password: password,
		}

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := tataiskClient.CreateFilm(ctx, &pbTataisk.CreateFilmRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func Regapp(name string, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		req := models.RegappRequest{
			Name:   name,
			Secret: secret,
		}

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := ssoClient.Regapp(ctx, &pbSSO.RegappRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func Login(email string, password string, appID int32) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		req := models.LoginRequest{
			Email:    email,
			Password: password,
			AppID:    appID,
		}

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := ssoClient.Login(ctx, &pbSSO.LoginRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func CreateFilm(data models.FilmData) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		req := models.FilmData{
			Title:        data.Title,
			YearOfProd:   data.YearOfProd,
			Imdb:         data.Imdb,
			Description:  data.Description,
			Country:      data.Country,
			Genre:        data.Genre,
			FilmDirector: data.FilmDirector,
			Screenwriter: data.Screenwriter,
			Budget:       data.Budget,
			Collection:   data.Collection,
		}

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := tataiskClient.CreateFilm(ctx, &pbTataisk.CreateFilmRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)

	}
}

func ReadFilm(id int32) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		resp, err := tataiskClient.ReadFilm(ctx, &pbTataisk.ReadFilmRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, resp)
	}
}

func UpdateFilm(id int32, data models.FilmData) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		req := models.FilmData{
			Title:        data.Title,
			YearOfProd:   data.YearOfProd,
			Imdb:         data.Imdb,
			Description:  data.Description,
			Country:      data.Country,
			Genre:        data.Genre,
			FilmDirector: data.FilmDirector,
			Screenwriter: data.Screenwriter,
			Budget:       data.Budget,
			Collection:   data.Collection,
		}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := tataiskClient.UpdateFilm(ctx, &pbTataisk.UpdateFilmRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func DeleteFilm(id int32) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		resp, err := tataiskClient.DeleteFilm(ctx, &pbTataisk.DeleteFilmRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp)
	}
}
