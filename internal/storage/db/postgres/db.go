package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	connStr := "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s"
	connStr = fmt.Sprintf(connStr, cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	// fmt.Printf("DB: %s\n", connStr)

	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertToTable(db *sqlx.DB, query string, params ...interface{}) (interface{}, error) {

	var result interface{}

	row := db.QueryRow(query, params...)
	if err := row.Scan(&result); err != nil {
		return &result, err
	}

	return result, nil
}
