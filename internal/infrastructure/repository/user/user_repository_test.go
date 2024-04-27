package user

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE users (
		id TEXT PRIMARY KEY,
		user_type_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		cpf TEXT NOT NULL,
		cnpj TEXT DEFAULT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted INTEGER NOT NULL DEFAULT 0,
		UNIQUE (cpf),
		UNIQUE (email),
		UNIQUE (cnpj)
		)`)
	if err != nil {
		return err
	}
	return nil
}

func createTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestCreateUser(t *testing.T) {
	db, err := createTestDB()
	if err != nil {
		t.Fatalf("error create database: %v", err)
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		t.Fatalf("error create table: %v", err)
	}

	userRepository := NewUserRepository(db)
	user, err := entity.CreateUserFactory(
		nil,
		"André Luiz",
		"andreluizmicro@gmail.com",
		"1z#sddAA2ttt",
		"088.022.254.80",
		nil,
		1,
	)

	_, err = userRepository.Create(user)
	assert.Nil(t, err)
}
