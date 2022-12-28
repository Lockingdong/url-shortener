package repository

import (
	"UrlShortener/internal/domain/entity"
	"context"
)

type ICreateUrlCodeRepository interface {
	SaveUrlInfo(ctx context.Context, ent *entity.UrlInfo) error
}

type IGetUrlRepository interface {
	GetUrlInfo(ctx context.Context, urlCode string) (*entity.UrlInfo, error)
}

type IUrlInfoRepository interface {
	ICreateUrlCodeRepository
	IGetUrlRepository
}
