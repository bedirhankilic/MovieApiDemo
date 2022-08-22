package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	businesslayer "github.com/bedirhankilic/movieapicase/business-layer"
	"github.com/bedirhankilic/movieapicase/constants"
	_ "github.com/bedirhankilic/movieapicase/docs"
	responsemodels "github.com/bedirhankilic/movieapicase/models/response-models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var HttpRoute = gin.New()

func getRoutes() {
	HttpRoute.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	HttpRoute.POST("/Login", Login)
	HttpRoute.GET("/GetAllMovies", GetAllMovies)

	authGroup := HttpRoute.Group("/")
	authGroup.Use(Auth())
	{
		authGroup.PUT("/AddMovie", AddMovie)
		authGroup.POST("/DeleteMovie/:movieId", DeleteMovie)
		authGroup.PUT("/UpdateMovie", UpdateMovie)
	}

	HttpRoute.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func InitHttpServer() {
	getRoutes()
	HttpRoute.Run(":5000")
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {

		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.JSON(401, responsemodels.GenericResponse{
				ResponseCode: constants.REQUIRED_LOGIN,
			})
			context.Abort()
			return
		}

		userService := businesslayer.NewAuthService()
		err := userService.ValidateToken(tokenString)
		if err != constants.OK {
			context.JSON(401, responsemodels.GenericResponse{
				ResponseCode: constants.JWT_TOKEN_EXPERIED,
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
