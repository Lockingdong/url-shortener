package application

import (
	"context"
)

type IGetOGUrlUseCase interface {
	Execute(ctx context.Context, qry *GetOGUrlUseCaseQuery) (*GetOGUrlUseCaseResponse, error)
}

type GetOGUrlUseCaseQuery struct {
	ShortUrlCode string `form:"short_url_code" binding:"required"`
}

type GetOGUrlUseCaseResponse struct {
	OGUrl string `json:"og_url"`
}

// TODO: implement "GetOGUrlUseCase"
