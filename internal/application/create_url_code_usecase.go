package application

import (
	"UrlShortener/internal/domain/entity"
	"UrlShortener/internal/domain/repository"
	"context"

	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

var _ ICreateUrlCodeUseCase = (*CreateUrlCodeUseCase)(nil)

type ICreateUrlCodeUseCase interface {
	Execute(ctx context.Context, cmd *CreateUrlCodeCommand) (*CreateUrlCodeResponse, error)
}

type CreateUrlCodeCommand struct {
	Url string `json:"url" binding:"required,uri"`
}

type CreateUrlCodeResponse struct {
	UrlCode string `json:"url_code"`
}

type CreateUrlCodeUseCase struct {
	repository repository.ICreateUrlCodeRepository
}

func NewCreateUrlCodeUseCase(
	repository repository.ICreateUrlCodeRepository,
) *CreateUrlCodeUseCase {
	return &CreateUrlCodeUseCase{
		repository: repository,
	}
}

func (u *CreateUrlCodeUseCase) Execute(
	ctx context.Context,
	cmd *CreateUrlCodeCommand,
) (*CreateUrlCodeResponse, error) {

	urlCode := shortid.MustGenerate()

	urlInfo := entity.NewUrlInfo(
		uuid.NewString(),
		cmd.Url,
		urlCode,
	)

	if err := u.repository.SaveUrlInfo(ctx, urlInfo); err != nil {
		return nil, err
	}

	return &CreateUrlCodeResponse{
		UrlCode: urlCode,
	}, nil
}
