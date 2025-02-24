package httpadapter

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vet-clinic-back/metrics-service/docs" // swagger
	"github.com/vet-clinic-back/metrics-service/internal/config"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"net/http"
	"time"
)

const httpTimeOut = 10 * time.Second

type HTTPAdapter struct {
	port       string
	origins    []string
	mainServer *http.Server
}

type HTTPHandler interface {
	GetMetrics(c *gin.Context)
}

func New(config config.HTTPConfig) *HTTPAdapter {
	return &HTTPAdapter{port: config.Port, origins: config.AllowOrigins}
}

func (s *HTTPAdapter) SetHandlers(httpHandler HTTPHandler) {
	gin.SetMode(gin.ReleaseMode)
	mainRouter := gin.New()
	mainRouter.GET("/metrics/", httpHandler.GetMetrics)
	mainRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if len(s.origins) > 0 {
		mainRouter.Use(cors.New(cors.Config{
			AllowOrigins: s.origins,
		}))
	}

	s.mainServer = &http.Server{
		Addr:         ":" + s.port,
		Handler:      mainRouter,
		WriteTimeout: httpTimeOut,
		ReadTimeout:  httpTimeOut,
	}
}

func (s *HTTPAdapter) MustRun() {
	log := logging.GetLogger().WithField("op", "HTTPServer.MustRun")

	log.WithField("port", s.port).Info("metrics http server started")
	err := s.mainServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithError(err).Fatal("metrics http server error")
	}
	if err != nil {
		log.Info("metrics http server closed")
	}
}
