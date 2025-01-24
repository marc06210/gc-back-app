package todo_test

import (
	"github.com/marc06210/gc-back-app/internal/model"
	"github.com/marc06210/gc-back-app/internal/todo"
	"reflect"
	"testing"
)

type MockDb struct {
}

func (m MockDb) GetAllPublications() ([]model.Publication, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockDb) Close() {
	//TODO implement me
	panic("implement me")
}

func TestService_Search(t *testing.T) {
	tests := []struct {
		testName string
		todos    []string
		query    string
		want     []string
	}{
		{
			testName: "Search",
			todos:    []string{"shopping"},
			query:    "sh",
			want:     []string{"shopping"},
		},
		{
			testName: "Search case insensitive",
			todos:    []string{"shopping"},
			query:    "Sh",
			want:     []string{"shopping"},
		},
	}
	mockDb := &MockDb{}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			svc := todo.NewService(mockDb)
			for _, todoItem := range tt.todos {
				_ = svc.Add(todoItem)
			}
			if got := svc.Search(tt.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
