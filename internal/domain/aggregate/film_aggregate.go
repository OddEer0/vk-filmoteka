package aggregate

import (
	appValidator "github.com/OddEer0/ck-filmoteka/internal/common/lib/app_validator"
	"github.com/OddEer0/ck-filmoteka/internal/domain/model"
)

type FilmAggregate struct {
	Film   model.Film
	Actors []*model.Actor
}

func (f *FilmAggregate) Validation() error {
	validator := appValidator.New()
	err := validator.Struct(f.Film)
	if err != nil {
		return err
	}
	return nil
}

func NewFilmAggregate(film model.Film) (*FilmAggregate, error) {
	result := &FilmAggregate{Film: film}
	if err := result.Validation(); err != nil {
		return nil, err
	}
	return result, nil
}
