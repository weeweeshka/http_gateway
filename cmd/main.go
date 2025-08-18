package main

import (
	"github.com/weeweeshka/http_gateway/internal/lib/logger"
	"github.com/weeweeshka/http_gateway/internal/router"
)

func main() {
	logr := logger.SetupLogger()
	r := router.SetupRouter()
	r.Run(":8080")
	logr.Info("HTTPGateway started on port :8080")
}
