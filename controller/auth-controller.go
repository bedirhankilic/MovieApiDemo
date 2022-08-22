package controller

import (
	"net/http"

	businesslayer "github.com/bedirhankilic/movieapicase/business-layer"
	"github.com/bedirhankilic/movieapicase/constants"
	requestmodels "github.com/bedirhankilic/movieapicase/models/request-models"
	responsemodels "github.com/bedirhankilic/movieapicase/models/response-models"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      This required for login
// @Description  This required for login
// @Tags         login
// @Accept       json
// @Produce      json
// @Param loginModel body requestmodels.LoginModel{} true "requestmodels.LoginModel Login item"
// @Success      200  {object}  responsemodels.TokenResponse{}
// @Failure      400  {object}  responsemodels.GenericResponse{}
// @Failure      401  {object}  responsemodels.GenericResponse{}
// @Router       /Login/ [post]
func Login(ctx *gin.Context) {
	var login requestmodels.LoginModel
	if ctx.ShouldBindJSON(&login) != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responsemodels.GenericResponse{
			ResponseCode: constants.WRONG_REQ_FORMAT,
		})
		return
	}

	userService := businesslayer.NewAuthService()
	user, errCode := userService.Login(login.Username, login.Password)
	if errCode != constants.OK {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responsemodels.GenericResponse{
			ResponseCode: constants.NOT_FOUND,
		})
		return
	}
	token, errJWT := userService.GenerateJWT(user.Email, user.Username)
	if errJWT != constants.OK {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responsemodels.GenericResponse{
			ResponseCode: constants.NOT_FOUND,
		})
		return
	}
	ctx.JSON(http.StatusOK, responsemodels.TokenResponse{
		GenericResponse: responsemodels.GenericResponse{
			ResponseCode: constants.NOT_FOUND,
		},
		Data: token,
	})
}
