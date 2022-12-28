package application

import (
	"context"
)

type IGetUrlUseCase interface {
	Execute(ctx context.Context, qry *GetUrlUseCaseQuery) (*GetUrlUseCaseResponse, error)
}

type GetUrlUseCaseQuery struct {
	UrlCode string `form:"url_code" binding:"required"`
}

type GetUrlUseCaseResponse struct {
	Url string `json:"url"`
}

// TODO: implement "GetUrlUseCase"
