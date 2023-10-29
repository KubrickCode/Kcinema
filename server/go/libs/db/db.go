package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB struct {
	Gorm *gorm.DB
}

func New() (*DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres"
	return open(dsn)
}

func open(dsn string) (*DB, error) {
	pg := postgres.Open(dsn)

	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	}

	gdb, err := gorm.Open(pg, cfg)
	if err != nil {
		return nil, err
	}

	return &DB{
		Gorm:	gdb,
	}, nil
}