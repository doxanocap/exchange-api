package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	ExchangersKeysTable       = "exchangers_keys"
	ExchangersCurrenciesTable = "exchangers_currencies"
	ExchangersInfoTable       = "exchangers_info"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func InitDB(cfg Config) *sqlx.DB {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))

	if err != nil {
		logrus.Fatalf("Error in postgres DB init: %s", err.Error())
		return nil
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatalf("Error in postgres DB init: %s", err.Error())
		return nil
	}
	return db
}
