package input_port

import (
	"UrlShortener/internal/application"
	"UrlShortener/internal/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlShortenerController struct {
	createUrlCodeUseCase application.ICreateUrlCodeUseCase
	getUrlUseCase        application.IGetUrlUseCase
}

func NewUrlShortenerController(
	repository repository.IUrlInfoRepository,
) *UrlShortenerController {
	return &UrlShortenerController{
		createUrlCodeUseCase: application.NewCreateUrlCodeUseCase(repository),
	}
}

func (c *UrlShortenerController) CreateUrlCode(ctx *gin.Context) {

	cmd := &application.CreateUrlCodeCommand{}
	if err := ctx.ShouldBind(cmd); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	result, err := c.createUrlCodeUseCase.Execute(ctx, cmd)
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

func (c *UrlShortenerController) GetUrl(ctx *gin.Context) {

	qry := &application.GetUrlUseCaseQuery{}
	if err := ctx.ShouldBind(qry); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	result, err := c.getUrlUseCase.Execute(ctx, qry)
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
