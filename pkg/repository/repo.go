package repository

type UrlList interface {
	GetUrl(shortUrl string) (string, error)
	GetShortUrl(longUrl string) (string, error)
	PostUrl(shortUrl string, longUrl string) error
	IsAvailable(shortUrl string) (bool, error)
}

type Repo struct {
	UrlList
}

func NewRepo(db UrlList) *Repo {
	return &Repo{
		UrlList: db,
	}
}
