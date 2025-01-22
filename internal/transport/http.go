package transport

import (
	"encoding/json"
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
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/publications", func(writer http.ResponseWriter, request *http.Request) {
		publications, err := todoSvc.GetAllPublications()
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(writer).Encode(publications)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /todo", func(writer http.ResponseWriter, request *http.Request) {

		b, err2 := json.Marshal(todoSvc.GetAll())
		if err2 != nil {
			log.Println(err2)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err := writer.Write(b)
		if err != nil {
			log.Println(err2)
			writer.WriteHeader(http.StatusInternalServerError)

		}
		writer.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		err = todoSvc.Add(t.Item)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusCreated)
		return
	})

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}