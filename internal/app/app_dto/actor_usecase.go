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

	GetByQueryResult struct {
		Actors    []*aggregate.ActorAggregate
		PageCount int
	}
)
