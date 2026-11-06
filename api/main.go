package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err == nil {
		defer conn.Close(context.Background())
	}
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte(uuid.NewString())) })
	_ = http.ListenAndServe(":8080", r)
}
