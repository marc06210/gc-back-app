package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/marc06210/gc-back-app/internal/model"
	"github.com/marc06210/gc-back-app/internal/publication"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type TodoItem struct {
	Item string `json:"item"`
}

type PublicationItem struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Host        string    `json:"host"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Creationts  time.Time `json:"creationts"`
}

type Server struct {
	router *gin.Engine
	logger *zap.Logger
}

// This handler extract the error from any following handler
// prints it and then return a 500 HTTP error code with
// no additionnal information
func errorHandler(c *gin.Context, logger *zap.Logger) {
	c.Next()

	for _, err := range c.Errors {
		// log, handle, etc.
		logger.Error("Error detected: %s\n", zap.Error(err))
	}

	c.JSON(http.StatusInternalServerError, "")
}
func zapLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log details after request processing
		logger.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("latency", time.Since(start)),
		)
	}
}

func postArticle(logger *zap.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		var publicationItem PublicationItem
		err := context.BindJSON(&publicationItem)
		if err != nil {
			return
		}
		now := time.Now()
		publicationItem.Creationts = now

		publicationIcon, err := model.IconOf(publicationItem.Icon)
		if err != nil {
			context.Error(err)
			return
		}

		logger.Debug("Creating publication", zap.Any("publication ", publicationItem), zap.String("icon", publicationIcon.String()))
		context.JSON(http.StatusNotImplemented, "")
	}
}

func NewServer(todoSvc *publication.Service, logger *zap.Logger) *Server {

	router := gin.Default()

	// two ways to define handlers injecting the logger
	router.Use(func(context *gin.Context) {
		errorHandler(context, logger)
	})
	router.Use(zapLoggerMiddleware(logger))

	router.GET("/api/publications", func(context *gin.Context) {
		logger.Debug("Getting all publications")
		publications, err := todoSvc.GetAllPublications()
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, publications)
		if err != nil {
			context.Error(err)
			return
		}
	})

	router.POST("/api/publications", postArticle(logger))

	return &Server{router: router, logger: logger}
}

func (s *Server) Serve() error {
	s.logger.Info("Starting server")
	return s.router.Run(":8080")
}
