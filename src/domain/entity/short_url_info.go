package entity

import "time"

type ShortUrlInfo struct {
	ID           string
	OGUrl        string
	ShortUrlCode string
	CreatedAt    time.Time
}

func NewShortUrlInfo(id, ogUrl, shortUrlCode string) *ShortUrlInfo {
	return &ShortUrlInfo{
		ID:           id,
		OGUrl:        ogUrl,
		ShortUrlCode: shortUrlCode,
		CreatedAt:    time.Now(),
	}
}
