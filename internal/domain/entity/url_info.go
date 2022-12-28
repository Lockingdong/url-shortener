package entity

import (
	"time"

	"github.com/google/uuid"
)

type UrlInfoParams struct {
	CreatedAt *time.Time
	ID        string
	Url       string
	UrlCode   string
}

type UrlInfo struct {
	CreatedAt time.Time
	ID        string
	Url       string
	UrlCode   string
}

func NewUrlInfo(params *UrlInfoParams) *UrlInfo {
	if params.ID == "" {
		params.ID = uuid.NewString()
	}

	if params.CreatedAt == nil {
		createdAt := time.Now()
		params.CreatedAt = &createdAt
	}

	return &UrlInfo{
		ID:        params.ID,
		Url:       params.Url,
		UrlCode:   params.UrlCode,
		CreatedAt: *params.CreatedAt,
	}
}
