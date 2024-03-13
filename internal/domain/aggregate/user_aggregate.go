package aggregate

import "github.com/OddEer0/ck-filmoteka/internal/domain/model"

type UserAggregate struct {
	User  model.User
	Token *model.Token
}
