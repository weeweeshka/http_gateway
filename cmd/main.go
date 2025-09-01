package main

import (
	"github.com/weeweeshka/http_gateway/internal/lib/logger"
	"github.com/weeweeshka/http_gateway/internal/router"
)

// @title Tataisk API
// @version 1.0
// @description API для работы с пользователями и фильмами
// @host localhost:8080
// @BasePath /
func main() {
	logr := logger.SetupLogger()
	r := router.SetupRouter()
	r.Run(":8080")
	logr.Info("HTTPGateway started on port :8080")
}
