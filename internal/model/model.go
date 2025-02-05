package model

import (
	"errors"
	"time"
)

// my custom type for the icons
// in database it is saved as an int
type Icon int64
const (
	SPRING Icon = 0
	ISTIO      	= 1
	ECS        	= 2
	UNK 		= -99
)

// this methods translate the enum value into a string representation
// used to feed the model
func (s Icon) String() string {
	switch s {
	case SPRING:
		return "spring"
	case ISTIO:
		return "istio"
	case ECS:
		return "ecs"
	}
	return "unknown"
}

func IconOf(representation string) (Icon, error) {
	switch representation {
	case "spring": return SPRING, nil
	case "istio": return ISTIO, nil
	case "ecs": return ECS, nil
	}
	return UNK, errors.New("unknown representation")
}

type Publication struct {
	Id          int64     `db:"id" json:"id"`
	Description string    `db:"description" json:"description"`
	Icon        string    `db:"icon" json:"icon"`
	Host        string    `db:"host" json:"host"`
	Title       string    `db:"title" json:"title"`
	Url         string    `db:"url" json:"url"`
	Creationts  time.Time `db:"creationts" json:"creationts"`
}
