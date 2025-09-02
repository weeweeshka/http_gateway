package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	pbComments "github.com/weeweeshka/comments_proto/gen/go/comments"
	"github.com/weeweeshka/http_gateway/internal/domain/models"
	clients "github.com/weeweeshka/http_gateway/internal/grpcClients"
	pbSSO "github.com/weeweeshka/sso_proto/gen/go/sso"
	pbTataisk "github.com/weeweeshka/tataisk_proto/gen/go/tataisk"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
)

var (
	ssoClient, _, _      = clients.SetupGateway()
	_, tataiskClient, _  = clients.SetupGateway()
	_, _, commentsClient = clients.SetupGateway()
)

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var req models.RegRequest

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := ssoClient.Register(ctx, &pbSSO.RegisterRequest{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func Regapp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		var req models.RegappRequest

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := ssoClient.Regapp(ctx, &pbSSO.RegappRequest{
			Name:   req.Name,
			Secret: req.Secret,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		var req models.LoginRequest

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := ssoClient.Login(ctx, &pbSSO.LoginRequest{
			Email:    req.Email,
			Password: req.Password,
			AppId:    req.AppID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func CreateFilm() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		var req models.FilmData
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"authorization": token,
		}))

		resp, err := tataiskClient.CreateFilm(ctx, &pbTataisk.CreateFilmRequest{
			Title:        req.Title,
			YearOfProd:   req.YearOfProd,
			Imdb:         req.Imdb,
			Description:  req.Description,
			Country:      req.Country,
			Genre:        req.Genre,
			FilmDirector: req.FilmDirector,
			Screenwriter: req.Screenwriter,
			Budget:       req.Budget,
			Collection:   req.Collection,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)

	}
}

func ReadFilm() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		filmIdStr := c.Param("id")
		filmID, _ := strconv.Atoi(filmIdStr)

		ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"authorization": token,
		}))

		resp, err := tataiskClient.ReadFilm(ctx, &pbTataisk.ReadFilmRequest{
			Id: int32(filmID),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, resp)
	}
}

func UpdateFilm() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		var req models.FilmData
		filmIdStr := c.Param("id")
		filmID, _ := strconv.Atoi(filmIdStr)

		ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"authorization": token,
		}))
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := tataiskClient.UpdateFilm(ctx, &pbTataisk.UpdateFilmRequest{
			Id:           int32(filmID),
			Title:        &req.Title,
			YearOfProd:   &req.YearOfProd,
			Imdb:         &req.Imdb,
			Description:  &req.Description,
			Country:      req.Country,
			Genre:        req.Genre,
			FilmDirector: &req.FilmDirector,
			Screenwriter: &req.Screenwriter,
			Budget:       &req.Budget,
			Collection:   &req.Collection,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func DeleteFilm() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		filmIdStr := c.Param("id")
		filmID, _ := strconv.Atoi(filmIdStr)

		ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"authorization": token,
		}))
		resp, err := tataiskClient.DeleteFilm(ctx, &pbTataisk.DeleteFilmRequest{
			Id: int32(filmID),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp)
	}
}

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		var req models.CommentReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		comment := &pbComments.Comment{
			Title:   req.Title,
			Content: req.Content,
		}

		ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"authorization": token,
		}))

		resp, err := commentsClient.CreateComment(ctx, &pbComments.CreateCommentRequest{FilmId: req.FilmID, Comment: comment})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func ReadComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		filmIdStr := c.Param("id")
		filmID, _ := strconv.Atoi(filmIdStr)

		ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"authorization": token,
		}))

		resp, err := commentsClient.GetComments(ctx, &pbComments.GetCommentsRequest{
			FilmId: int32(filmID),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		filmIdStr := c.Param("id")
		filmID, _ := strconv.Atoi(filmIdStr)
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"authorization": token,
		}))

		resp, err := commentsClient.DeleteComment(ctx, &pbComments.DeleteCommentRequest{
			FilmId: int32(filmID),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, resp)

	}
}
