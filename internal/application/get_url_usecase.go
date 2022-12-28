package application

import (
	"UrlShortener/internal/domain/repository"
	"context"
)

var _ IGetUrlUseCase = (*GetUrlUseCase)(nil)

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
type GetUrlUseCase struct {
	repository repository.IGetUrlRepository
}

func NewGetUrlUseCase(
	repository repository.IGetUrlRepository,
) *GetUrlUseCase {
	return &GetUrlUseCase{
		repository: repository,
	}
}

func (u *GetUrlUseCase) Execute(ctx context.Context, qry *GetUrlUseCaseQuery) (*GetUrlUseCaseResponse, error) {

	urlInfo, err := u.repository.GetUrlInfo(ctx, qry.UrlCode)
	if err != nil {
		return nil, err
	}

	return &GetUrlUseCaseResponse{
		Url: urlInfo.Url,
	}, nil
}
