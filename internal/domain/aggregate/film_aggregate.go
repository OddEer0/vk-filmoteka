package aggregate

import "github.com/OddEer0/ck-filmoteka/internal/domain/model"

type FilmAggregate struct {
	Film   model.Film
	Actors []*model.Actor
}

func (f *FilmAggregate) Validation() error {
	return nil
}

func NewFilmAggregate(film model.Film) (*FilmAggregate, error) {
	result := &FilmAggregate{Film: film}
	if err := result.Validation(); err != nil {
		return nil, err
	}
	return result, nil
}
