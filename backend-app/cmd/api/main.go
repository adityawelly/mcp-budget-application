package main

import (
	"backend/cmd/api/models"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var conn config

	flag.IntVar(&conn.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&conn.env, "env", "development", "Application environment (developement|production)")
	flag.StringVar(&conn.db.dsn, "dsn", "postgres://postgres:password@localhost/mcp_budget_app?sslmode=disable", "Postgres connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(conn)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		config: conn,
		logger: logger,
		models: models.NewModels(db),
	}

	// fmt.Println("Jalan")

	// http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "masuk")
	// 	currentStatus := AppStatus{
	// 		Status:      "Available",
	// 		Environment: conn.env,
	// 		Version:     version,
	// 	}

	// 	js, err := json.MarshalIndent(currentStatus, "", "\t")
	// 	if err != nil {
	// 		log.Println(err)
	// 	}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(js)
	// })

	// err := http.ListenAndServe(fmt.Sprintf(":%d", conn.port), nil)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", conn.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", conn.port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}

func openDB(conn config) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
