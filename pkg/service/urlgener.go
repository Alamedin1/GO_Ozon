package service

import (
	"math/rand"
	"sample-app/pkg/repository"
)

type UrlGener struct {
	repo   repository.UrlList
	strLen int
	_rune  []rune
}

func NewUrlGener(repo repository.UrlList, strLen int, _rune []rune) *UrlGener {
	chars := make([]rune, 0, len(_rune))
	chars = append(chars, _rune...)
	return &UrlGener{
		repo:   repo,
		strLen: strLen,
		_rune:  chars,
	}
}

func (u *UrlGener) UrlGet(shortUrl string) (string, error) {
	return u.repo.UrlGet(shortUrl)
}

func (u *UrlGener) UrlPost(longUrl string) (string, error) {
	var shortUrl string
	var err error
	shortUrlAvailable := false
	for !shortUrlAvailable {
		shortUrl = u.getUniqueString()
		shortUrlAvailable, err = u.repo.IsAvailable(shortUrl)
		if err != nil {
			return "", err
		}
	}
	err = u.repo.UrlPost(shortUrl, longUrl)
	if err != nil {
		shortUrl, errQuery := u.repo.ShortUrlGet(longUrl)
		if errQuery != nil {
			return "", errQuery
		}
		return shortUrl, err
	}
	return shortUrl, nil
}

func (u *UrlGener) getUniqueString() string {
	uniqueRuneArray := make([]rune, u.strLen)
	for i := range uniqueRuneArray {
		uniqueRuneArray[i] = u._rune[rand.Intn(len(u._rune))]
	}
	return string(uniqueRuneArray)
}
