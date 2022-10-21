package service

import "sample-app/pkg/repository"

type UrlList interface {
	UrlGet(shortUrl string) (string, error)
	UrlPost(longUrl string) (string, error)
}

type Service struct {
	UrlList
}

func NewService(repo *repository.Repository, strLen int, _rune []rune) *Service {
	return &Service{
		UrlList: NewUrlGener(repo.UrlList, strLen, _rune),
	}
}
