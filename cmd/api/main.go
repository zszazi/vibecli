package main

import (
	"log"

	"github.com/joho/godotenv"
	"zszazi.github.io/vibecli/internal/db"
	"zszazi.github.io/vibecli/internal/env"
	"zszazi.github.io/vibecli/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env variables")
	}
	cfg := config{
		addr: env.GetString("VIBECLI_ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("VIBECLI_DB_ADDR", "postgres://admin:adminpassword@localhost/vibecli?sslmode=disable"),
			maxOpenConns: env.GetInt("VIBECLI_DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("VIBECLI_DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetTimeDuration("VIBECLI_DB_MAX_IDLE_TIME", "10m"),
		},
		env:     env.GetString("VIBECLI_ENV", "dev"),
		version: env.GetString("VIBECLI_VERSION", ""),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")
	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
