package main

import "database/sql"

type UserStore interface {
	RegisterUser(u RegisterPayload, db *sql.DB) error
}

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{
		db: db,
	}
}
