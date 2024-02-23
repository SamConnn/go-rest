package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

// CreateAccount implements Storage.
func (*PostgresStorage) CreateAccount(*Account) error {
	panic("unimplemented")
}

// DeleteAccount implements Storage.
func (*PostgresStorage) DeleteAccount(id int) error {
	panic("unimplemented")
}

// GetAccountByID implements Storage.
func (*PostgresStorage) GetAccountByID(id int) (*Account, error) {
	panic("unimplemented")
}

// UpdateAccount implements Storage.
func (*PostgresStorage) UpdateAccount(*Account) error {
	panic("unimplemented")
}

func NewPostgresStorage() (*PostgresStorage, error) {
	db, err := sql.Open("postgres", "host=172.22.246.106 dbname=default user=admin password=admin@DEV1 sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}
