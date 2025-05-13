package database

import (
	"errors"
	"log/slog"

	"github.com/joyharbour/joyharbour-server/server/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDatabase() error {
	dialector, err := loadDatabaseDialector()
	if err != nil {
		return err
	}

	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func loadDatabaseDialector() (gorm.Dialector, error) {
	config := config.GetConfig().Database
	slog.Info("init database", "type", config.Type)

	if config.Type == "sqlite" {
		return sqlite.Open(config.Dsn), nil
	} else if config.Type == "mysql" {
		return mysql.Open(config.Dsn), nil
	} else if config.Type == "postgres" {
		return postgres.Open(config.Dsn), nil
	} else if config.Type == "sqlserver" {
		return sqlserver.Open(config.Dsn), nil
	}

	return nil, errors.New("unknown database type")
}
