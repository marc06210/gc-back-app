package publication

import (
	"errors"
	"github.com/marc06210/gc-back-app/internal/db"
	"github.com/marc06210/gc-back-app/internal/model"
	"log"
	"strings"
)

type Item struct {
	Task   string
	Status string
}

type Service struct {
	todos []Item
	db    db.Interface
}

func NewService(db db.Interface) *Service {
	return &Service{
		todos: make([]Item, 0),
		db:    db,
	}
}

func (svc *Service) Add(todo string) error {
	for _, t := range svc.todos {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}
	svc.todos = append(svc.todos, Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	})
	return nil
}

func (svc *Service) GetAll() []Item {
	return svc.todos
}

func (svc *Service) GetAllPublications() ([]model.Publication, error) {
	data, err := svc.db.GetAllPublications()
	if err != nil {
		return nil, err
	}
	log.Printf("obtained %d publications", len(data))
	return data, nil
}

func (svc *Service) Search(query string) []string {
	var results []string
	for _, item := range svc.todos {
		if strings.Contains(strings.ToLower(item.Task), strings.ToLower(query)) {
			results = append(results, item.Task)
		}
	}
	return results
}
