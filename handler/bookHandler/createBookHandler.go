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

// @Summary Create book
// @Description Create book
// @Tags Book
// @Accept json
// @Produce json
// @Param request body CreateBookRequest true "Request body"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /book/create [post]
func CreateBookHandler (ctx *gin.Context) {


	request := CreateBookRequest{}

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


	
	timestamp, err := time.Parse(handler.CustomLayout, request.ReleaseDate)
	book := schemas.Book{
		Title: request.Title,
		Pages: request.Pages,
		ReleaseDate: timestamp,
		AuthorID: uint(request.AuthorID),
		Author: author,
	}

	err = handler.Db.Create(&book).Error


	if err != nil {
		handler.Logger.ErrorF("error creating book: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError,"error creating book")
		return
	}

	handler.SendSuccess(ctx, "create-book", book)




}