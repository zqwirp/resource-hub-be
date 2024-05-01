package database

import (
	"database/sql"
	"reshub/pkg/database"
)

type SQLConnector interface {
	GetDB() *sql.DB
}

type SQL struct {
	db *sql.DB
}

func NewSQL() SQLConnector {
	db, _ := database.OpenSQL("postgres", "onta")
	return &SQL{db: db}
}

func (db *SQL) GetDB() *sql.DB {
	return db.db
}
