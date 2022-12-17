package application

import (
	"UrlShortener/internal/domain/entity"
	"UrlShortener/internal/domain/repository"
	"context"

	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

var _ ICreateShortUrlCodeUseCase = (*CreateShortUrlCodeUseCase)(nil)

type ICreateShortUrlCodeUseCase interface {
	Execute(ctx context.Context, cmd *CreateShortUrlCodeCommand) (*CreateShortUrlCodeResponse, error)
}

type CreateShortUrlCodeCommand struct {
	OGUrl string `json:"og_url"`
}

type CreateShortUrlCodeResponse struct {
	ShortUrlCode string `json:"short_url_code"`
}

type CreateShortUrlCodeUseCase struct {
	repository repository.ICreateShortUrlCodeRepository
}

func NewCreateShortUrlUseCase(
	repository repository.ICreateShortUrlCodeRepository,
) *CreateShortUrlCodeUseCase {
	return &CreateShortUrlCodeUseCase{
		repository: repository,
	}
}

func (u *CreateShortUrlCodeUseCase) Execute(
	ctx context.Context,
	cmd *CreateShortUrlCodeCommand,
) (*CreateShortUrlCodeResponse, error) {

	shortUrlCode := shortid.MustGenerate()

	shortUrlInfo := entity.NewShortUrlInfo(
		uuid.NewString(),
		cmd.OGUrl,
		shortUrlCode,
	)

	if err := u.repository.CreateShortUrlInfo(ctx, shortUrlInfo); err != nil {
		return nil, err
	}

	return &CreateShortUrlCodeResponse{
		ShortUrlCode: shortUrlCode,
	}, nil
}
