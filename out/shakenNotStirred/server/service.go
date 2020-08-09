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

// custom enum type definitions

type BeverageType string

const (
	Beer BeverageType = "Beer"
	Liquor BeverageType = "Liquor"
	Wine BeverageType = "Wine"
)

type UserInput struct{
}
type BeverageInput struct{
	Name string
	Proof int
	Type BeverageType
}

[{"Name":"user","Description":"","Fields":[],"Namespaces":null},{"Name":"beverage","Description":"","Fields":[{"Name":"name","Description":"","TypeOf":"string","Primitive":{"Nullable":false,"MaxLength":0},"Selector":true,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null},{"Name":"proof","Description":"","TypeOf":"number","Primitive":{"Nullable":false,"DisableNegativeValues":false,"EnableDecimal":false,"MaxValue":0},"Selector":true,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null},{"Name":"type","Description":"","TypeOf":"beverageType","Primitive":{"Nullable":false,"Name":"beverageType","Values":["BEER","LIQUOR","WINE"]},"Selector":true,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null}],"Namespaces":null}]
[{"Name":"user","Description":"","Fields":[{"Name":"firstName","Description":"","TypeOf":"name","Primitive":{"Nullable":false,"MaxLength":64},"Selector":false,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null},{"Name":"lastName","Description":"","TypeOf":"name","Primitive":{"Nullable":false,"MaxLength":64},"Selector":false,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null},{"Name":"joinedAt","Description":"","TypeOf":"datetime","Primitive":{},"Selector":false,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null},{"Name":"phoneNumber","Description":"","TypeOf":"phoneNumber","Primitive":{"Nullable":false,"MaxLength":10},"Selector":false,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null}],"Namespaces":null},{"Name":"cocktail","Description":"","Fields":[{"Name":"name","Description":"","TypeOf":"string","Primitive":{"Nullable":false,"MaxLength":0},"Selector":true,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null}],"Namespaces":null},{"Name":"beverage","Description":"","Fields":[{"Name":"name","Description":"","TypeOf":"string","Primitive":{"Nullable":false,"MaxLength":0},"Selector":true,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null},{"Name":"proof","Description":"","TypeOf":"number","Primitive":{"Nullable":false,"DisableNegativeValues":false,"EnableDecimal":false,"MaxValue":0},"Selector":true,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null},{"Name":"type","Description":"","TypeOf":"beverageType","Primitive":{"Nullable":false,"Name":"beverageType","Values":["BEER","LIQUOR","WINE"]},"Selector":true,"SelectorTypes":null,"Optional":false,"DefaultValue":null,"Namespaces":null}],"Namespaces":null}]
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
type Beverage struct{
	ID uuid.UUID
	Archived bool
	Name string
	Proof int
	Type BeverageType
}

const beverageCreateQueryStr = `
INSERT INTO beverage (
	"id",
	"name",
	"proof",
	"type"
) VALUES (
	$1,
	$2,
	$3,
	$4
);
`

func (s *ServiceInstance) CreateBeverage() (Beverage, error) {
	input := Beverage{
		ID: uuid.New(),
		Archived: false,
		Name: "",
		Proof: 0,
		Type: "",
	}
	
	err := s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(beverageCreateQueryStr, input.Name, input.Proof, input.Type)
		return err
	})
	if err != nil {
		return Beverage{}, err
	}
	
	return Beverage{}, nil
}

func (s *ServiceInstance) ArchiveBeverage(id string) error {
	return s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec("UPDATE beverage SET \"_archived\" = TRUE WHERE \"id\" = $1", id)
		return err
	})
}