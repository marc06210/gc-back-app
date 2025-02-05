package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/marc06210/gc-back-app/internal/model"
	"github.com/marc06210/gc-back-app/internal/todo"
	"log"
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
}

// This handler extract the error from any following handler
// prints it and then return a 500 HTTP error code with
// no additionnal information
func errorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		// log, handle, etc.
		log.Printf("Error detected: %s\n", err)
	}

	c.JSON(http.StatusInternalServerError, "")
}

func postArticle(context *gin.Context) {
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

	log.Printf("Publication to create: %s with %d icon\n", publicationItem, publicationIcon)
	context.JSON(http.StatusNotImplemented, "")
}

func NewServer(todoSvc *todo.Service) *Server {

	router := gin.Default()

	router.Use(errorHandler)

	router.GET("/api/publications", func(context *gin.Context) {
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

	router.POST("/api/publications", postArticle)

	return &Server{router: router}
}

func (s *Server) Serve() error {
	return s.router.Run(":8080")
}
