package main

import (
	"time"

	"github.com/google/uuid"
)

type ServiceInstance struct{
	DB string
}

type User {
	ID uuid.UUID
	FirstName string
	LastName string
	JoinedAt time.Time
	PhoneNumber string
}

const userCreateQueryStr = `
INSERT INTO user (
	"firstName",
	"lastName",
	"joinedAt",
	"phoneNumber"
) VALUES (
	$1,
	$2,
	$3,
	$4
);
`

func (s *ServiceInstance) CreateUser() (User, error) {
	err := s.DB.Update(func(tx *sqlx.Tx) error {
		_, err := tx.Exec(userCreateQueryStr, input.firstName, input.lastName, input.joinedAt, input.phoneNumber)
		return err
	})
	if err != nil {
		return User{}, err
	}
	
	return User{}, nil
}

func ArchiveUser(id string) error {
	return nil
}

func GetUser(id string) (User, error) {
	return User{}, nil
}

func UpdateUser(id string) (User, error) {
	return User{}, nil
}


type Cocktail {
	ID uuid.UUID
	Name string
}

const cocktailCreateQueryStr = `
INSERT INTO cocktail (
	"name"
) VALUES (
	$1
);
`

func (s *ServiceInstance) CreateCocktail() (Cocktail, error) {
	err := s.DB.Update(func(tx *sqlx.Tx) error {
		_, err := tx.Exec(cocktailCreateQueryStr, input.name)
		return err
	})
	if err != nil {
		return Cocktail{}, err
	}
	
	return Cocktail{}, nil
}

func ArchiveCocktail(id string) error {
	return nil
}

func GetCocktail(id string) (Cocktail, error) {
	return Cocktail{}, nil
}

func UpdateCocktail(id string) (Cocktail, error) {
	return Cocktail{}, nil
}


