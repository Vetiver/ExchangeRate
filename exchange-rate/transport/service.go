package transport

import (
	"exchangeRate/transport/handlers"
	"exchangeRate/pkg"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpServer struct {
	logger   *zap.Logger
	httpPort string
}

func NewHttpServer(logger *zap.Logger, httpPort string) *HttpServer {
	return &HttpServer{
		logger:   logger,
		httpPort: httpPort,
	}
}

func (h *HttpServer) StartHTTPServer(handlers *handlers.BaseHandler) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(pkg.AccessLog())

	router.GET("/ping", func(c *gin.Context) {
		handlers.Ping(c)
	})

	router.GET("/currency", func(c *gin.Context) {
		handlers.ExchangeRate(c)
	})


	h.logger.Info("HTTP server is running on port", zap.String("port", h.httpPort))
	if err := router.Run(h.httpPort); err != nil {
		h.logger.Fatal("Failed to start HTTP server", zap.Error(err))
	}
}
