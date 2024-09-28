package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type DBService struct {
	db *pgxpool.Pool
}

func NewDBService() *DBService {
	return &DBService{db: CreatePool("connection_string")}
}

func CreatePool(constr string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(constr)

	if err != nil {
		log.Fatal(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}

func (db *DBService) GetAllUsers() ([]string, error) {
	return nil, nil
}
