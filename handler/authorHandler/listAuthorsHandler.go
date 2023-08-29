package authorHandler

import (
	"net/http"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary List authors
// @Description List authors
// @Tags Author
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /author/list [get]
func ListAuthorsHandler (ctx *gin.Context) {

	authors := []schemas.Author{}


	err := handler.Db.Find(&authors).Error

	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "error listing authors")
		return
	}

	handler.SendSuccess(ctx, "list-authors", authors)

}