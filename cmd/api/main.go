package main

import (
	"context"
	"log"
	"net/http"
	"time"

	httpx "example.com/goprac11-borisovda/internal/http"
	"example.com/goprac11-borisovda/internal/http/handlers"
	"example.com/goprac11-borisovda/internal/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	databaseURL := "postgres://user123:pass123@localhost:5432/notes_db?sslmode=disable"

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("Failed to parse DB config: ", err)
	}

	cfg.MaxConns = 20
	cfg.MinConns = 5
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute
	cfg.ConnConfig.ConnectTimeout = 10 * time.Second

	log.Printf("Connecting to PostgreSQL at %s...", databaseURL)

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal("Failed to create connection pool: ", err)
	}
	defer pool.Close()

	repo := repo.NewNoteRepoPG(pool)
	h := &handlers.Handler{Repo: repo}
	r := httpx.NewRouter(h)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
