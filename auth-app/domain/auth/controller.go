package auth

import (
	"net/http"
	"tirenn/test-efishery/auth-app/dto/http/request"
	"tirenn/test-efishery/auth-app/dto/http/response"
	"tirenn/test-efishery/auth-app/utils"

	"github.com/gin-gonic/gin"
)

type ControllerContract interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetJWT(ctx *gin.Context)
}

type Controller struct {
	service ServiceContract
}

func NewController(service ServiceContract) ControllerContract {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Create(ctx *gin.Context) {
	var req request.Auth
	var res response.Response

	if err := ctx.BindJSON(&req); err != nil {
		res.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := utils.ValidateStruct(req); err != nil {
		res.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data, err := c.service.Create(req)
	if err != nil {
		res.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Success"
	res.Data = data
	ctx.JSON(http.StatusCreated, res)
}

func (c *Controller) Login(ctx *gin.Context) {
	var req request.Login
	var res response.Response

	if err := ctx.BindJSON(&req); err != nil {
		res.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := utils.ValidateStruct(req); err != nil {
		res.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data, err := c.service.Login(req)
	if err != nil {
		res.Message = "Unauthorized"
		ctx.JSON(http.StatusUnauthorized, res)
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
