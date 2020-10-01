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

null
[{"Name":"user","Description":"","Fields":[{"Name":"lastName","Description":"","TypeOf":"name","Primitive":{"Nullable":false,"MaxLength":64},"Selector":false,"SelectorTypes":[],"Optional":false,"DefaultValue":null,"Namespaces":[]},{"Name":"lastName","Description":"","TypeOf":"name","Primitive":{"Nullable":false,"MaxLength":64},"Selector":false,"SelectorTypes":[],"Optional":false,"DefaultValue":null,"Namespaces":[]},{"Name":"joinedAt","Description":"","TypeOf":"datetime","Primitive":{},"Selector":false,"SelectorTypes":[],"Optional":false,"DefaultValue":null,"Namespaces":[]},{"Name":"phoneNumber","Description":"","TypeOf":"phoneNumber","Primitive":{"Nullable":false,"MaxLength":10},"Selector":false,"SelectorTypes":[],"Optional":false,"DefaultValue":null,"Namespaces":[]}],"Namespaces":[]}]
type User struct{
	ID uuid.UUID
	LastName string
	LastName string
	JoinedAt time.Time
	PhoneNumber string
}

const userCreateQueryStr = `
INSERT INTO user (
	"id",
	"lastName",
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
		LastName: "",
		LastName: "",
		JoinedAt: time.Time{},
		PhoneNumber: "",
	}
	
	err := s.DB.Update(s.Context, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(userCreateQueryStr, input.LastName, input.LastName, input.JoinedAt, input.PhoneNumber)
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