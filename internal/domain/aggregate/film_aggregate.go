package aggregate

import "github.com/OddEer0/ck-filmoteka/internal/domain/model"

type FilmAggregate struct {
	Film   model.Film
	Actors []*model.Actor
}
