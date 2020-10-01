package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"context"
	"os"

	"github.com/cosmotek/pgdb"
	"github.com/rs/zerolog"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"goji.io"
	"goji.io/pat"
)

func main() {
	logger := zerolog.New(os.Stdout)
	db, err := pgdb.Dial(pgdb.Config{
		User: "user",
		Password: "password",
		Host: "localhost",
		Port: "5432",
		DatabaseName: "shakenNotStirred",
		MigrationDir: "./migrations",
		SSLDisabled: true,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
		MaxConnLifespan: time.Second * time.Duration(30),
	})
	if err != nil {
		panic(err)
	}

	migrationsDiff, err := db.DiffMigrations()
	if err != nil {
		panic(err)
	}

	currentMigration, err := db.GetCurrentMigration()
	if err != nil {
		panic(err)
	}

	migrationStatus, err := db.RunMigrations(logger, currentMigration, migrationsDiff...)
	if err != nil {
		panic(err)
	}

	log.Printf("applied %v\n", migrationStatus.Applied)
	log.Printf("skipped %v\n", migrationStatus.Skipped)
	log.Printf("failed %v\n", migrationStatus.Failed)
	log.Printf("latest %v\n", migrationStatus.Latest)

	service := ServiceInstance{DB: db, Context: context.Background()}
	mux := goji.NewMux()

	mux.HandleFunc(pat.Options("/rpc/v1/createUser"), func(res http.ResponseWriter, req *http.Request) {
		input := UserInput{}
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			http.Error(res, fmt.Sprintf("createShakenNotStirred failed to decode input: '%s'", err.Error()), http.StatusBadRequest)
			return
		}

		created, err := service.CreateUser(input)
		if err != nil {
			hErr, ok := err.(HTTPError)
			if ok {
				http.Error(res, hErr.Error(), hErr.Code())
				return
			}

			http.Error(res, err.Error(), 500)
			return
		}

		// todo filter out fields by namespace
		// todo use generated JSON stubs for perf improvements

		err = json.NewEncoder(res).Encode(created)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		go eventsteam.Push(Event{
			TypeOf: UserCreated,
			ResourceID: created.ID,
			CompletedAt: time.Now().UTC(),
		})
	})
	mux.HandleFunc(pat.Options("/rpc/v1/deleteUser"), func(res http.ResponseWriter, req *http.Request) {
		params := struct{ ID string }{}
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		
		err = service.DeleteUser(params.ID)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	})
	log.Println("starting service on port 5000...")
	http.ListenAndServe(":5000", mux)
}