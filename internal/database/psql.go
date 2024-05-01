package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PSQLConnector interface {
	GetDB() *pgx.Conn
}

type PSQL struct {
	db *pgx.Conn
}

func Something() {
	connString := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greeting)
}
