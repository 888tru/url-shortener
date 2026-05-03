package service

import "github.com/888tru/url-shortener/internal/repository"

type URLServiceImpl struct {
	repo repository.URLRepository
}

func NewUrlService(repo repository.URLRepository) *URLServiceImpl {
	return &URLServiceImpl{repo: repo}
}

func (s *URLServiceImpl) CreateShortUrl(origUrl string) (string, error) {
	return "", nil
}

func (s *URLServiceImpl) GetOriginalUrl(code string) (string, error) {
	return "", nil
}
