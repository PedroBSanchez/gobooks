package bookhandler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Update book
// @Description Update book
// @Tags Book
// @Accept json
// @Produce json
// @Param request body UpdateBookRequest true "Request body"
// @Param id query string true "Book identification"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /book/update [put]
func UpdateBookHandler (ctx *gin.Context) {

	bookId := ctx.Query("id")


	if bookId == "" {
		handler.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	request := UpdateBookRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		handler.Logger.ErrorF("validation error: %v, %v", err.Error(), request.Pages)
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}


	author := schemas.Author{};


	err = handler.Db.First(&author, request.AuthorID).Error


	if err != nil {
		handler.Logger.ErrorF("author with id: %v not found", request.AuthorID)
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("author with id: %v not found", request.AuthorID))
		return
	}


	book := schemas.Book{}

	err = handler.Db.First(&book, bookId).Error
	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("book with id: %s not found", bookId))
		return 
	}


	timestamp, err := time.Parse(handler.CustomLayout, request.ReleaseDate)
	book.Author = author
	book.AuthorID = author.ID
	book.Title = request.Title
	book.Pages = request.Pages
	book.ReleaseDate = timestamp


	err = handler.Db.Save(&book).Error


	if err != nil {
		handler.Logger.ErrorF("error updating book: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError,"error updating book")
		return
	}

	handler.SendSuccess(ctx, "update-book", book)




}