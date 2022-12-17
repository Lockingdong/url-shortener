package output_port

import (
	"UrlShortener/internal/domain/entity"
	"UrlShortener/internal/domain/repository"
	"context"
	"fmt"
	"time"
)

var _ repository.IShortUrlInfoRepository = (*ShortUrlInfoMockRepository)(nil)

type ShortUrlInfoMockDTO struct {
	ID           string
	OGUrl        string
	ShortUrlCode string
	CreatedAt    time.Time
}

type ShortUrlInfoMockRepository struct {
	shortUrlInfos map[string]*ShortUrlInfoMockDTO
}

func NewShortUrlInfoMockRepository(
	shortUrlInfos map[string]*ShortUrlInfoMockDTO,
) *ShortUrlInfoMockRepository {

	if shortUrlInfos == nil {
		shortUrlInfos = map[string]*ShortUrlInfoMockDTO{}
	}

	return &ShortUrlInfoMockRepository{
		shortUrlInfos: shortUrlInfos,
	}
}

func (r *ShortUrlInfoMockRepository) CreateShortUrlInfo(
	_ context.Context,
	ent *entity.ShortUrlInfo,
) error {

	dto := &ShortUrlInfoMockDTO{
		ID:           ent.ID,
		OGUrl:        ent.OGUrl,
		ShortUrlCode: ent.ShortUrlCode,
		CreatedAt:    ent.CreatedAt,
	}

	r.shortUrlInfos[dto.ShortUrlCode] = dto

	return nil
}

func (r *ShortUrlInfoMockRepository) GetShortUrlInfo(
	_ context.Context,
	shortUrlCode string,
) (*entity.ShortUrlInfo, error) {

	if val, ok := r.shortUrlInfos[shortUrlCode]; ok {

		ent := &entity.ShortUrlInfo{
			ID:           val.ID,
			OGUrl:        val.OGUrl,
			ShortUrlCode: val.ShortUrlCode,
			CreatedAt:    val.CreatedAt,
		}

		return ent, nil
	}

	return nil, fmt.Errorf("[GetShortUrlInfo] short url code %s not found", shortUrlCode)
}
