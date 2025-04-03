package repository

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(ctx context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UsersStore{db: db},
	}
}
