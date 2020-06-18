package main

import (
	"encoding/json"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	service := ServiceInstance{}
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
	var dUser = `
	/rpc/v1/archiveUser
	`
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
	var dCocktail = `
	/rpc/v1/archiveCocktail
	`
}
