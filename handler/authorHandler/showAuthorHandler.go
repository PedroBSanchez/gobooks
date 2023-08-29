package authorHandler

import (
	"fmt"
	"net/http"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Show author
// @Description Show an author
// @Tags Author
// @Accept json
// @Produce json
// @Param id query string true "Author identification"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /author/show [get]
func ShowAuthorHandler(ctx *gin.Context) {

	authorId := ctx.Query("id")


	if authorId == "" {
	   handler.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	   return
	}

	author := schemas.Author{}


	err := handler.Db.First(&author, authorId).Error

	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("author with id: %s not found", authorId))
		return 
	}
 
	handler.SendSuccess(ctx, "show-author", author)

}