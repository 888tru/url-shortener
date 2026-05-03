package service

import "context"

type URLService interface {
	CreateShortUrl(ctx context.Context, originalUrl string) (string, error)
	GetOriginalUrl(ctx context.Context, code string) (string, error)
}
