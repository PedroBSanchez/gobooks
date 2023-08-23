package router

import (
	docs "github.com/PedroBSanchez/gobooks.git/docs"
	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {


	handler.InitializeHandler()

	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath

	author := router.Group(basePath + "/author")
	{
		author.GET("/")
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}