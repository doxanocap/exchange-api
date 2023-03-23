package repository

import "github.com/jmoiron/sqlx"

type PoolModel struct {
	psql *sqlx.DB
}

func NewPoolModel(db *sqlx.DB) *PoolModel {
	return &PoolModel{psql: db}
}
