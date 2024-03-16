package inMemDb

import "github.com/OddEer0/vk-filmoteka/internal/domain/model"

type ActorFilm struct {
	ActorId string
	FilmId  string
}

type InMemDb struct {
	Users     []*model.User
	Tokens    []*model.Token
	Actor     []*model.Actor
	Film      []*model.Film
	ActorFilm []*ActorFilm
}

var instance *InMemDb = nil

func New() *InMemDb {
	if instance != nil {
		return instance
	}

	instance = &InMemDb{
		Users:     []*model.User{},
		Tokens:    []*model.Token{},
		Actor:     []*model.Actor{},
		Film:      []*model.Film{},
		ActorFilm: []*ActorFilm{},
	}

	return instance
}
