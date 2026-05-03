package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) Save(ctx context.Context, code string, originalUrl string) error {
	_, err := p.db.Exec(ctx, "insert into urls (code, original_url) values ($1, $2)", code, originalUrl)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepository) Get(ctx context.Context, code string) (string, error) {
	var originalUrl string
	err := p.db.QueryRow(ctx, "select original_url from urls where code = $1", code).Scan(&originalUrl)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
