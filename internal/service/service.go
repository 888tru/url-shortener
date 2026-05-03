package service

type URLService interface {
	CreateShortUrl(originalUrl string) (string, error)
	GetOriginalUrl(code string) (string, error)
}
