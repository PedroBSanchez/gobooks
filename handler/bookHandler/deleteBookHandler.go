package bookhandler

import (
	"fmt"
	"net/http"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)


func DeleteBookHandler(ctx *gin.Context) {

	bookId := ctx.Query("id")

	if bookId == "" {
		handler.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	book := schemas.Book{}


	err := handler.Db.First(&book, bookId).Error


	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("book with id: %s not found", bookId))
		return
	}


	//Delete book
	err = handler.Db.Delete(&book).Error

	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting book with id: %s", bookId))
		return
	}


	handler.SendSuccess(ctx, "delete-book", book)


}