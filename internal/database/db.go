package database

import (
	"context"
	"database/sql"
	"log"
	"reshub/config"
	"reshub/pkg/database"

	"github.com/jackc/pgx/v5"
)

type Closer interface {
	Close() error
}

type SQL struct {
	db *sql.DB
}

type PSQL struct {
	conn *pgx.Conn
}

func NewSQL(cfg *config.Config) Closer {
	db, err := database.OpenSQL("postgres", "onta")
	if err != nil {
		log.Fatal(err)
	}

	return &SQL{db: db}
}

func (d *SQL) Close() error {
	return d.db.Close()
}

func NewPSQL(cfg *config.Config) Closer {
	conn, err := pgx.Connect(context.Background(), "onta")
	if err != nil {
		log.Fatal(err)
	}
	return &PSQL{conn: conn}
}

func (d *PSQL) Close() error {
	return d.conn.Close(context.Background())
}

// func Something() {
// 	connString := "user=pqgotest dbname=pqgotest sslmode=verify-full"
// 	conn, err := pgx.Connect(context.Background(), connString)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(context.Background())

// 	var greeting string
// 	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(greeting)
// }
