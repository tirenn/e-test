package fetch

import (
	"net/http"
	"tirenn/test-efishery/fetch-app/dto/http/response"

	"github.com/gin-gonic/gin"
)

type ControllerContract interface {
	Get(ctx *gin.Context)
	GetJWT(ctx *gin.Context)
	GetAggregate(ctx *gin.Context)
}

type Controller struct {
	service ServiceContract
}

func NewController(service ServiceContract) ControllerContract {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Get(ctx *gin.Context) {
	var res response.Response

	data, err := c.service.Get()
	if err != nil {
		res.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Success"
	res.Data = data
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) GetJWT(ctx *gin.Context) {
	var res response.Response

	token := ctx.Param("token")

	if token == "" {
		res.Message = "Token Required"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data, err := c.service.GetJWT(token)
	if err != nil {
		res.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Success"
	res.Data = data
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) GetAggregate(ctx *gin.Context) {
	var res response.Response

	data, err := c.service.GetAggregate()
	if err != nil {
		res.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Success"
	res.Data = data
	ctx.JSON(http.StatusOK, res)
}
