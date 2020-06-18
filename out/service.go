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
	Archived bool
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
		Archived: false,
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

func (s *ServiceInstance) ArchiveUser(id string) error {
	return s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec("UPDATE user SET \"_archived\" = TRUE WHERE \"id\" = $1", id)
		return err
	})
}
type Cocktail struct{
	ID uuid.UUID
	Archived bool
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
		Archived: false,
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

func (s *ServiceInstance) ArchiveCocktail(id string) error {
	return s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec("UPDATE cocktail SET \"_archived\" = TRUE WHERE \"id\" = $1", id)
		return err
	})
}
