package entity

import "time"

type UrlInfo struct {
	CreatedAt time.Time
	ID        string
	Url       string
	UrlCode   string
}

func NewUrlInfo(id, url, urlCode string) *UrlInfo {
	return &UrlInfo{
		ID:        id,
		Url:       url,
		UrlCode:   urlCode,
		CreatedAt: time.Now(),
	}
}
