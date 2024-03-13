package aggregate

import "github.com/OddEer0/ck-filmoteka/internal/domain/model"

type ActorAggregate struct {
	Actor model.Actor
	Films []*model.Film
}
