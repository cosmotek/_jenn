package main

import (
	"context"
	"time"

	"github.com/cosmotek/pgdb"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ServiceInstance struct{
	Context context.Context
	DB *pgdb.Database
}

type User struct{
	ID uuid.UUID
	FirstName string
	LastName string
	JoinedAt time.Time
	PhoneNumber string
}

const userCreateQueryStr = `
INSERT INTO user (
	"id",
	"firstName",
	"lastName",
	"joinedAt",
	"phoneNumber"
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
);
`

func (s *ServiceInstance) CreateUser() (User, error) {
	input := User{
		ID: uuid.New(),
		FirstName: "",
		LastName: "",
		JoinedAt: time.Time{},
		PhoneNumber: "",
	}
	
	err := s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(userCreateQueryStr, input.FirstName, input.LastName, input.JoinedAt, input.PhoneNumber)
		return err
	})
	if err != nil {
		return User{}, err
	}
	
	return User{}, nil
}

func (s *ServiceInstance) DeleteUser(id string) error {
	return s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec("DELETE FROM user WHERE \"id\" = $1", id)
		return err
	})
}

func GetUser(id string) (User, error) {
	return User{}, nil
}

func UpdateUser(id string) (User, error) {
	return User{}, nil
}


type Cocktail struct{
	ID uuid.UUID
	Name string
}

const cocktailCreateQueryStr = `
INSERT INTO cocktail (
	"id",
	"name"
) VALUES (
	$1,
	$2
);
`

func (s *ServiceInstance) CreateCocktail() (Cocktail, error) {
	input := Cocktail{
		ID: uuid.New(),
		Name: "",
	}
	
	err := s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(cocktailCreateQueryStr, input.Name)
		return err
	})
	if err != nil {
		return Cocktail{}, err
	}
	
	return Cocktail{}, nil
}

func (s *ServiceInstance) DeleteCocktail(id string) error {
	return s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec("DELETE FROM cocktail WHERE \"id\" = $1", id)
		return err
	})
}

func GetCocktail(id string) (Cocktail, error) {
	return Cocktail{}, nil
}

func UpdateCocktail(id string) (Cocktail, error) {
	return Cocktail{}, nil
}


