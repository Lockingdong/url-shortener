package repository

import (
	"UrlShortener/internal/domain/entity"
	"context"
)

type ICreateShortUrlCodeRepository interface {
	SaveShortUrlInfo(ctx context.Context, ent *entity.ShortUrlInfo) error
}

type IGetOGUrlRepository interface {
	GetShortUrlInfo(ctx context.Context, shortUrlCode string) (*entity.ShortUrlInfo, error)
}

type IShortUrlInfoRepository interface {
	ICreateShortUrlCodeRepository
	IGetOGUrlRepository
}
