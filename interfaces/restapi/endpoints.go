package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"

	"github.com/cosmotek/_jenn/app"
)

func registerRoutes() http.Handler {
	mux := goji.NewMux()

	mux.HandleFunc(pat.Get("/version"), func(res http.ResponseWriter, req *http.Request) {
		version := app.Version()
		err := json.NewEncoder(res).Encode(map[string]string{
			"jenn_version": version,
		})
		if err != nil {
			http.Error(res, fmt.Sprintf("failed to serialize response: '%s'", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc(pat.Get("/version"), func(res http.ResponseWriter, req *http.Request) {
		version := app.Version()
		err := json.NewEncoder(res).Encode(map[string]string{
			"jenn_version": version,
		})
		if err != nil {
			http.Error(res, fmt.Sprintf("failed to serialize response: '%s'", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	return mux
}

func Start(ctxWithCancel context.Context) error {
	server := http.Server{
		Addr:    "",
		Handler: registerRoutes(),
	}

	go func() {
		<-ctxWithCancel.Done()

		// attempt graceful shutdown first
		server.Shutdown(ctxWithCancel)
		server.Close()
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
