package repository

import "context"

type URLRepository interface {
	Save(ctx context.Context, code string, originalUrl string) error
	Get(ctx context.Context, code string) (string, error)
}
