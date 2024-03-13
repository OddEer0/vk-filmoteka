package aggregate

import "github.com/OddEer0/ck-filmoteka/internal/domain/model"

type ActorAggregate struct {
	Actor model.Actor
	Films []*model.Film
}

func (a *ActorAggregate) Validation() error {
	return nil
}

func NewActorAggregate(actor model.Actor) (*ActorAggregate, error) {
	result := &ActorAggregate{Actor: actor}
	if err := result.Validation(); err != nil {
		return nil, err
	}
	return result, nil
}
