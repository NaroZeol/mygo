package controller

import (
	"mygo/internal/pkg/common"
	"mygo/internal/pkg/constants"
	"mygo/internal/pkg/utils"
	"mygo/internal/server/service"

	"github.com/gin-gonic/gin"
)

// @Summary		login
// @Description	login
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			loginRequest	body		common.LoginRequest	true	"login request"
// @Success		200				{object}	common.Result		"OK"
// @Router			/api/login [post]
func Login(ctx *gin.Context) {
	var loginRequest common.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	username, password := loginRequest.Username, loginRequest.Password

	id, role, err := service.Login(username, password)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}

	token, err := utils.GenerateToken(id, username, role)
	if err != nil {
		ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
		return
	}

	ctx.Header(constants.TOKEN_NAME, constants.TOKEN_PREFIX+token)
	ctx.JSON(common.OK, common.Ok(nil))
}

// @Summary		register
// @Description	register
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			registerRequest	body		common.RegisterRequest	true	"register request"
// @Success		200				{object}	common.Result			"OK"
// @Router			/api/register [post]
func Register(ctx *gin.Context) {
	var registerRequest common.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}
	username, password, role := registerRequest.Username, registerRequest.Password, registerRequest.Role

	id, err := service.Register(username, password, role)
	if err != nil {
		if common.CheckInternalError(err) {
			ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
			return
		}
		ctx.JSON(common.BAD_REQUEST, common.Bad(err.Error()))
		return
	}

	token, err := utils.GenerateToken(id, username, role)
	if err != nil {
		ctx.JSON(common.INTERNAL_SERVER_ERROR, common.InternalError(err.Error()))
		return
	}

	ctx.Header(constants.TOKEN_NAME, constants.TOKEN_PREFIX+token)
	ctx.JSON(common.OK, common.Ok(nil))
}