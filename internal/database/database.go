package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reshub/pkg/database"

	"github.com/jackc/pgx/v5"
)

type DBConnector interface {
	GetDB() *sql.DB
}

type SQL struct {
	db *sql.DB
}

func NewDatabase() DBConnector {
	db, _ := database.OpenSQL("postgres", "onta")
	return &SQL{db: db}
}

func (db *SQL) GetDB() *sql.DB {
	return nil
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
