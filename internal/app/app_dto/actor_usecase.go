package appDto

import (
	"github.com/OddEer0/vk-filmoteka/internal/domain/aggregate"
	"time"
)

type (
	CreateActorUseCaseDto struct {
		Name     string
		Gender   string
		Birthday time.Time
	}

	ActorGetByQueryResult struct {
		Actors    []*aggregate.ActorAggregate `json:"actors"`
		PageCount int                         `json:"pageCount"`
	}
)
