package bookhandler

import (
	"net/http"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary List books
// @Description List books
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /book/list [get]
func ListBookHandler (ctx *gin.Context) {

	books := []schemas.Book{}

	err := handler.Db.Preload("Author").Find(&books).Error


	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "error listing books")
		return
	}


	handler.SendSuccess(ctx,"list-books", books)

}