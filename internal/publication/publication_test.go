package publication_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/marc06210/gc-back-app/internal/model"
	"github.com/marc06210/gc-back-app/internal/publication"
	"testing"
	"time"
)

type MockDb struct {
}

func (m MockDb) GetAllPublications() ([]model.Publication, error) {
	return []model.Publication{
		{
			Id:          202,
			Description: "Logs are a major feature of an app, so learn how to create JUnit to also cover logs.",
			Icon:        "spring",
			Host:        "",
			Title:       "Asserting Log Messages With JUnit",
			Url:         "https://www.baeldung.com/junit-asserting-logs",
			Creationts:  time.Now(),
		},
		{
			Id:          203,
			Description: "In this tutorial, we will run a task hosted in AWS Fargate ECS anytime a message is published into a SQS queue. This solution is simpler than the solution presented in the previous article.",
			Icon:        "ecs",
			Host:        "",
			Title:       "Launch and run ECS task on a SQS publication - a simple solution",
			Url:         "https://medium.com/@marc.guerrini/aws-4-dummies-run-ecs-task-on-an-sqs-publication-a-simple-solution-f1b8567d4edf",
			Creationts:  time.Now(),
		},
	}, nil
}

func (m MockDb) Close() {
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
			svc := publication.NewService(mockDb)
			publications, _ := svc.GetAllPublications()
			assert.Equal(t, 2, len(publications))
		})
	}
}
