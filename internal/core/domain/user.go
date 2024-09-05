package domain

import (
	vo "github.com/axel-andrade/opina-ai-api/internal/core/domain/value_objects"
)

type User struct {
	Base
	Name     vo.Name     `json:"name"`
	Email    vo.Email    `json:"email"`
	Password vo.Password `json:"-"`
}

func BuildUser(name, email, password string) (*User, error) {
	u := &User{
		Name:     vo.Name{Value: name},
		Email:    vo.Email{Value: email},
		Password: vo.Password{Value: password},
	}

	if err := u.validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) validate() error {
	if err := u.Name.Validate(); err != nil {
		return err
	}

	if err := u.Email.Validate(); err != nil {
		return err
	}

	if err := u.Password.Validate(); err != nil {
		return err
	}

	return nil
}
