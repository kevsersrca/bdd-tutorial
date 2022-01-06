package user

import (
	"errors"
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

var (
	ErrInvalid = errors.New("invalid value")
	ErrConvert = errors.New("error convert string to int")
	ErrNull    = errors.New("null value")
)

type User interface {
	SetName(name string) error
	SetEmail(email string) error
	SetAge(age string) error
	GetName() string
	GetEmail() string
	GetAge() int
}

type user struct {
	name  string
	email string
	age   int
}

func NewUser() User {
	return &user{}
}

func (u *user) SetName(name string) error {
	if valid.IsNull(name) {
		return ErrNull
	}
	u.name = strings.ToUpper(name)
	return nil
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) SetEmail(email string) error {
	if valid.IsNull(email) {
		return ErrNull
	}

	if !valid.IsEmail(email) {
		return ErrInvalid
	}

	u.email = email
	return nil
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) SetAge(ageStr string) error {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return ErrConvert
	}

	if age <= 0 {
		return ErrInvalid
	}

	u.age = age
	return nil
}

func (u *user) GetAge() int {
	return u.age
}
