package authorHandler

import (
	"net/http"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Create author
// @Description Create a new job author
// @Tags Author
// @Accept json
// @Produce json
// @Param request body CreateAuthorRequest true "Request body"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /author [post]
func CreateAuthorHandler(ctx *gin.Context) {

	request := CreateAuthorRequest{}
	ctx.BindJSON(&request)

	err := request.Validate()
	if err != nil {
		handler.Logger.ErrorF("validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	author := schemas.Author{
		Name: request.Name,
		Age: request.Age,
		AmountBooks: 0,
	}


	err = handler.Db.Create(&author).Error

	if err != nil {
		handler.Logger.ErrorF("error creating author: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError,"error creating author")
		return
	}

	
	handler.SendSuccess(ctx, "create-author", author)

	

}





