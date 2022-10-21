package repository

type UrlList interface {
	UrlGet(shortUrl string) (string, error)
	ShortUrlGet(longUrl string) (string, error)
	UrlPost(shortUrl string, longUrl string) error
	IsAvailable(shortUrl string) (bool, error)
}

type Repository struct {
	UrlList
}

func NewRepo(db UrlList) *Repository {
	return &Repository{
		UrlList: db,
	}
}
