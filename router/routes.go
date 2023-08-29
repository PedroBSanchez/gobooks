package router

import (
	docs "github.com/PedroBSanchez/gobooks.git/docs"
	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/handler/authorHandler"
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
		author.POST("/create", authorHandler.CreateAuthorHandler)
		author.GET("/list", authorHandler.ListAuthorsHandler)
		author.GET("/show",authorHandler.ShowAuthorHandler)
		author.DELETE("/delete", authorHandler.DeleteAuthorHandler)
		author.PUT("/update", authorHandler.UpdateAuthorHandler)
	}

	book := router.Group(basePath + "/book")
	{
		book.GET("/")
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}