package main

import (
	"encoding/json"
	"net/http"
	"time"
	"context"

	"github.com/cosmotek/pgdb"

	"goji.io"
	"goji.io/pat"
)

func main() {
	db, err := pgdb.Dial(pgdb.Config{
		User: "user",
		Password: "password",
		Host: "localhost",
		Port: "5432",
		DatabaseName: "shakenNotStirred",
		SSLDisabled: true,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
		MaxConnLifespan: time.Second * time.Duration(30),
	})
	if err != nil {
		panic(err)
	}

	service := ServiceInstance{DB: db, Context: context.Background()}
	mux := goji.NewMux()

	mux.HandleFunc(pat.Options("/rpc/v1/createUser"), func(res http.ResponseWriter, req *http.Request) {
		created, err := service.CreateUser()
		if err != nil {
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
	})
	mux.HandleFunc(pat.Options("/rpc/v1/archiveUser"), func(res http.ResponseWriter, req *http.Request) {
		params := struct{ ID string }{}
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		
		err = service.ArchiveUser(params.ID)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	})
	mux.HandleFunc(pat.Options("/rpc/v1/createCocktail"), func(res http.ResponseWriter, req *http.Request) {
		created, err := service.CreateCocktail()
		if err != nil {
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
	})
	mux.HandleFunc(pat.Options("/rpc/v1/archiveCocktail"), func(res http.ResponseWriter, req *http.Request) {
		params := struct{ ID string }{}
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		
		err = service.ArchiveCocktail(params.ID)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	})
	http.ListenAndServe(":5000", mux)
}