package database

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
)

func OpenSQL(driverName, connStringParams string) (*sql.DB, error) {
	db, err := sql.Open(driverName, connStringParams)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

func PGX(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
