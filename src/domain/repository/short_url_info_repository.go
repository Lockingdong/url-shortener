package repository

import (
	"UrlShortener/src/domain/entity"
	"context"
)

type ICreateShortUrlCodeRepository interface {
	CreateShortUrlInfo(ctx context.Context, ent *entity.ShortUrlInfo) error
}

type IGetOGUrlRepository interface {
	GetShortUrlInfo(ctx context.Context, shortUrlCode string) (*entity.ShortUrlInfo, error)
}

type IShortUrlInfoRepository interface {
	ICreateShortUrlCodeRepository
	IGetOGUrlRepository
}
