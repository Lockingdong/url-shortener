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

var _ repository.IShortUrlInfoRepository = (*ShortUrlInfoRepository)(nil)

const keyPrefix string = "SHORT_URL"

type ShortUrlInfoDTO struct {
	ID           string    `json:"id"`
	OGUrl        string    `json:"og_url"`
	ShortUrlCode string    `json:"short_url_code"`
	CreatedAt    time.Time `json:"created_at"`
}

func (d *ShortUrlInfoDTO) ToJsonString() (string, error) {
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

type ShortUrlInfoRepository struct {
	cacheDBClient *redis.Client
}

func NewShortUrlInfoRepository(cacheDBClient *redis.Client) *ShortUrlInfoRepository {
	return &ShortUrlInfoRepository{
		cacheDBClient: cacheDBClient,
	}
}

func (r *ShortUrlInfoRepository) CreateShortUrlInfo(ctx context.Context, ent *entity.ShortUrlInfo) error {

	dto := &ShortUrlInfoDTO{
		ID:           ent.ID,
		OGUrl:        ent.OGUrl,
		ShortUrlCode: ent.ShortUrlCode,
		CreatedAt:    ent.CreatedAt,
	}

	key := getShortUrlKey(dto.ShortUrlCode)
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

func (r *ShortUrlInfoRepository) GetShortUrlInfo(ctx context.Context, shortUrlCode string) (*entity.ShortUrlInfo, error) {

	key := getShortUrlKey(shortUrlCode)
	jsonStr, err := r.cacheDBClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var dto ShortUrlInfoDTO
	if err := json.Unmarshal([]byte(jsonStr), &dto); err != nil {
		return nil, err
	}

	return &entity.ShortUrlInfo{
		ID:           dto.ID,
		OGUrl:        dto.OGUrl,
		ShortUrlCode: dto.ShortUrlCode,
		CreatedAt:    dto.CreatedAt,
	}, nil
}

func getShortUrlKey(code string) string {
	return keyPrefix + ":" + code
}
