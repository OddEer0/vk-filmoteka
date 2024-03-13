package aggregate

import (
	appValidator "github.com/OddEer0/ck-filmoteka/internal/common/lib/app_validator"
	"github.com/OddEer0/ck-filmoteka/internal/domain/model"
)

type UserAggregate struct {
	User  model.User
	Token *model.Token
}

func (u *UserAggregate) Validation() error {
	validator := appValidator.New()
	err := validator.Struct(u.User)
	if err != nil {
		return err
	}
	return nil
}

func NewUserAggregate(user model.User) (*UserAggregate, error) {
	result := &UserAggregate{User: user}
	if err := result.Validation(); err != nil {
		return nil, err
	}
	return result, nil
}
