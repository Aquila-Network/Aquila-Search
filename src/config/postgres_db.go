package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg *ConfigEnv) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Name,
		cfg.Database.Pass,
		cfg.Database.SslMode,
		// "db",
		// "5432",
		// "root",
		// "go_aquila",
		// "password",
		// "disable",
	))
	if err != nil {
		fmt.Println("Error postgres")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping")
		return nil, err
	}

	return db, nil
}
