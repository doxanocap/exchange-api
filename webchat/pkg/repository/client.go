package repository

import "github.com/jmoiron/sqlx"

type ClientModel struct {
	psql *sqlx.DB
}

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	City        string `json:"city"`
	PhoneNumber string `json:"phone_number"`
	Password    []byte `json:"-"`
}

func NewClientModel(db *sqlx.DB) *PoolModel {
	return &PoolModel{psql: db}
}
