package repository

import "github.com/jackc/pgx/v5/pgxpool"

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) Save(code string, originalUrl string) error {
	return nil
}

func (p *PostgresRepository) Get(code string) (string, error) {
	return "", nil
}
