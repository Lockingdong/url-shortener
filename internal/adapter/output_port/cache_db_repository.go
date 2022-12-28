package output_port

import (
	"UrlShortener/internal/domain/entity"
	"UrlShortener/internal/domain/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ repository.IUrlInfoRepository = (*UrlInfoRepository)(nil)

const keyPrefix string = "SHORT_URL"

type UrlInfoDTO struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	UrlCode   string    `json:"url_code"`
}

func (d *UrlInfoDTO) ToJsonString() (string, error) {
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

type UrlInfoRepository struct {
	cacheDBClient *redis.Client
}

func NewUrlInfoRepository(cacheDBClient *redis.Client) *UrlInfoRepository {
	return &UrlInfoRepository{
		cacheDBClient: cacheDBClient,
	}
}

func (r *UrlInfoRepository) SaveUrlInfo(ctx context.Context, ent *entity.UrlInfo) error {

	dto := &UrlInfoDTO{
		ID:        ent.ID,
		Url:       ent.Url,
		UrlCode:   ent.UrlCode,
		CreatedAt: ent.CreatedAt,
	}

	key := getUrlKey(dto.UrlCode)
	value, err := dto.ToJsonString()
	if err != nil {
		return err
	}

	res, err := r.cacheDBClient.SetNX(ctx, key, value, redis.KeepTTL).Result()
	if err != nil {
		return err
	}

	if res != true {
		return fmt.Errorf("Key already exist. key: %s", key)
	}

	return nil

}

func (r *UrlInfoRepository) GetUrlInfo(ctx context.Context, urlCode string) (*entity.UrlInfo, error) {

	key := getUrlKey(urlCode)
	jsonStr, err := r.cacheDBClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var dto UrlInfoDTO
	if err := json.Unmarshal([]byte(jsonStr), &dto); err != nil {
		return nil, err
	}

	return &entity.UrlInfo{
		ID:        dto.ID,
		Url:       dto.Url,
		UrlCode:   dto.UrlCode,
		CreatedAt: dto.CreatedAt,
	}, nil
}

func getUrlKey(code string) string {
	return keyPrefix + ":" + code
}
