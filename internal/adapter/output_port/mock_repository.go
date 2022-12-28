package output_port

import (
	"UrlShortener/internal/domain/entity"
	"UrlShortener/internal/domain/repository"
	"context"
	"fmt"
	"time"
)

var _ repository.IUrlInfoRepository = (*UrlInfoMockRepository)(nil)

type UrlInfoMockDTO struct {
	CreatedAt time.Time
	ID        string
	Url       string
	UrlCode   string
}

type UrlInfoMockRepository struct {
	urlInfos map[string]*UrlInfoMockDTO
}

func NewUrlInfoMockRepository(
	urlInfos map[string]*UrlInfoMockDTO,
) *UrlInfoMockRepository {

	if urlInfos == nil {
		urlInfos = map[string]*UrlInfoMockDTO{}
	}

	return &UrlInfoMockRepository{
		urlInfos: urlInfos,
	}
}

func (r *UrlInfoMockRepository) SaveUrlInfo(
	_ context.Context,
	ent *entity.UrlInfo,
) error {

	dto := &UrlInfoMockDTO{
		ID:        ent.ID,
		Url:       ent.Url,
		UrlCode:   ent.UrlCode,
		CreatedAt: ent.CreatedAt,
	}

	r.urlInfos[dto.UrlCode] = dto

	return nil
}

func (r *UrlInfoMockRepository) GetUrlInfo(
	_ context.Context,
	urlCode string,
) (*entity.UrlInfo, error) {

	if val, ok := r.urlInfos[urlCode]; ok {

		ent := entity.NewUrlInfo(&entity.UrlInfoParams{
			ID:        val.ID,
			Url:       val.Url,
			UrlCode:   val.UrlCode,
			CreatedAt: &val.CreatedAt,
		})

		return ent, nil
	}

	return nil, fmt.Errorf("[GetUrlInfo] url code %s not found", urlCode)
}
