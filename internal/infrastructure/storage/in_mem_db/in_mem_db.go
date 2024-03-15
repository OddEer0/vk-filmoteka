package inMemDb

import "github.com/OddEer0/ck-filmoteka/internal/domain/model"

type InMemDb struct {
	Users  []*model.User
	Tokens []*model.Token
	Actor  []*model.Actor
	Film   []*model.Film
}

var instance *InMemDb = nil

func New() *InMemDb {
	if instance != nil {
		return instance
	}

	instance = &InMemDb{
		Users:  []*model.User{},
		Tokens: []*model.Token{},
		Actor:  []*model.Actor{},
		Film:   []*model.Film{},
	}

	return instance
}
