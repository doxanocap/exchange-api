package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"os"
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

func Migrations(db *sqlx.DB) error {
	dest := ".\\pkg\\repository\\postgres"
	if len(os.Args) > 1 {
		if os.Args[1] == "up" {
			dest += "\\up.sql"
		}
		if os.Args[1] == "down" {
			dest += "\\down.sql"
		}

		content, err := os.ReadFile(dest)
		if err != nil {
			log.Fatal(err)
		}

		res, err := db.Query(string(content))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		res.Close()

		return errors.New(dest[1:])
	}
	return nil
}
