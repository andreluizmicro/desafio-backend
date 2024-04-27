package database

import (
	"database/sql"
	"fmt"
	"github.com/andreluizmicro/desafio-backend/config"
)

func NewConnection(cfg *config.AppConfig) (*sql.DB, error) {
	strConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open(cfg.DBDriver, strConnection)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
