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
		input := map[string]interface{}{}
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}

		e := json.NewEncoder(os.Stdout)
		e.SetIndent("", "\t")
		
		err = e.Encode(input)
		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}

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
	mux.HandleFunc(pat.Get("/rpc/v1/streamUser"), func(res http.ResponseWriter, req *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(req, res)
		if err != nil {
			// handle error
			http.Error(res, err.Error(), 500)
			return
		}

		go func() {
			defer conn.Close()

			ticker := time.NewTicker(time.Second * 3)
			defer ticker.Stop()

			for {
				select {
				case <-req.Context().Done():
					log.Println("closing stream.")
					return
				case <-ticker.C:
					msg, _, err := wsutil.ReadClientData(conn)
					if err != nil {
						// handle error
						http.Error(res, err.Error(), 500)
						return
					}
					log.Println(msg)

					err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(time.Now().String()))
					if err != nil {
						// handle error
						http.Error(res, err.Error(), 500)
						return
					}
				}
			}
		}()
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
		input := map[string]interface{}{}
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}

		e := json.NewEncoder(os.Stdout)
		e.SetIndent("", "\t")
		
		err = e.Encode(input)
		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}

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
	mux.HandleFunc(pat.Get("/rpc/v1/streamCocktail"), func(res http.ResponseWriter, req *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(req, res)
		if err != nil {
			// handle error
			http.Error(res, err.Error(), 500)
			return
		}

		go func() {
			defer conn.Close()

			ticker := time.NewTicker(time.Second * 3)
			defer ticker.Stop()

			for {
				select {
				case <-req.Context().Done():
					log.Println("closing stream.")
					return
				case <-ticker.C:
					msg, _, err := wsutil.ReadClientData(conn)
					if err != nil {
						// handle error
						http.Error(res, err.Error(), 500)
						return
					}
					log.Println(msg)

					err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(time.Now().String()))
					if err != nil {
						// handle error
						http.Error(res, err.Error(), 500)
						return
					}
				}
			}
		}()
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
	mux.HandleFunc(pat.Options("/rpc/v1/createBeverage"), func(res http.ResponseWriter, req *http.Request) {
		input := map[string]interface{}{}
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}

		e := json.NewEncoder(os.Stdout)
		e.SetIndent("", "\t")
		
		err = e.Encode(input)
		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}

		created, err := service.CreateBeverage()
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
	mux.HandleFunc(pat.Get("/rpc/v1/streamBeverage"), func(res http.ResponseWriter, req *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(req, res)
		if err != nil {
			// handle error
			http.Error(res, err.Error(), 500)
			return
		}

		go func() {
			defer conn.Close()

			ticker := time.NewTicker(time.Second * 3)
			defer ticker.Stop()

			for {
				select {
				case <-req.Context().Done():
					log.Println("closing stream.")
					return
				case <-ticker.C:
					msg, _, err := wsutil.ReadClientData(conn)
					if err != nil {
						// handle error
						http.Error(res, err.Error(), 500)
						return
					}
					log.Println(msg)

					err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(time.Now().String()))
					if err != nil {
						// handle error
						http.Error(res, err.Error(), 500)
						return
					}
				}
			}
		}()
	})
	mux.HandleFunc(pat.Options("/rpc/v1/archiveBeverage"), func(res http.ResponseWriter, req *http.Request) {
		params := struct{ ID string }{}
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		
		err = service.ArchiveBeverage(params.ID)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	})
	log.Println("starting service on port 5000...")
	http.ListenAndServe(":5000", mux)
}