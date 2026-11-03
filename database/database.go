package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID           int
	Usename      string
	Email        string
	PasswordHash string
}

func ConnectDB() *pgxpool.Pool {
	dsn := "postgres://rizki:rizki@localhost:5432/auth?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v\n", err)
	}

	fmt.Println("Koneksi database berhasil!")
	return pool
}
