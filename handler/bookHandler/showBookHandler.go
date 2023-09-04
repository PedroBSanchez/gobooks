package bookhandler

import (
	"fmt"
	"net/http"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Show book
// @Description Show a book
// @Tags Book
// @Accept json
// @Produce json
// @Param id query string true "Book identification"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /book/show [get]
func ShowBookHandler(ctx *gin.Context) {

	bookId := ctx.Query("id")


	if bookId == "" {
		handler.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	book := schemas.Book{}


	err := handler.Db.Preload("Author").First(&book, bookId).Error

	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("book with id: %s not found", bookId))
		return 
	}

	handler.SendSuccess(ctx, "show-book", book)

}