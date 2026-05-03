package service

import (
	"context"
	"math/rand"

	"github.com/888tru/url-shortener/internal/repository"
)

type URLServiceImpl struct {
	repo repository.URLRepository
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewUrlService(repo repository.URLRepository) *URLServiceImpl {
	return &URLServiceImpl{repo: repo}
}

func (s *URLServiceImpl) CreateShortUrl(ctx context.Context, origUrl string) (string, error) {
	shortCode := generateCode()
	if err := s.repo.Save(ctx, shortCode, origUrl); err != nil {
		return "", err
	}
	return shortCode, nil
}

func (s *URLServiceImpl) GetOriginalUrl(ctx context.Context, code string) (string, error) {
	origUrl, err := s.repo.Get(ctx, code)
	if err != nil {
		return "", err
	}
	return origUrl, nil
}

func generateCode() string {
	result := make([]byte, 6)
	for i := 0; i < 6; i++ {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(result)
}
