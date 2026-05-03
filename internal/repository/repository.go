package repository

type URLRepository interface {
	Save(code string, originalUrl string) error
	Get(code string) (string, error)
}
