package input_port

import (
	"UrlShortener/internal/application"
	"UrlShortener/internal/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortUrlController struct {
	createShortUrlUseCase application.ICreateShortUrlCodeUseCase
	getOGUrlUseCase       application.IGetOGUrlUseCase
}

func NewShortUrlInfoController(
	repository repository.IShortUrlInfoRepository,
) *ShortUrlController {
	return &ShortUrlController{
		createShortUrlUseCase: application.NewCreateShortUrlUseCase(repository),
	}
}

func (c *ShortUrlController) CreateShortUrlCode(ctx *gin.Context) {

	cmd := &application.CreateShortUrlCodeCommand{}
	if err := ctx.ShouldBind(cmd); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	result, err := c.createShortUrlUseCase.Execute(ctx, cmd)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		result,
	)
}

func (c *ShortUrlController) GetOGUrl(ctx *gin.Context) {

	qry := &application.GetOGUrlUseCaseQuery{}
	if err := ctx.ShouldBind(qry); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	result, err := c.getOGUrlUseCase.Execute(ctx, qry)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	ctx.PureJSON(
		http.StatusOK,
		result,
	)
}
