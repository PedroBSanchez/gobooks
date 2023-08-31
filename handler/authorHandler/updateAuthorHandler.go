package authorHandler

import (
	"fmt"
	"net/http"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Update author
// @Description Update an author
// @Tags Author
// @Accept json
// @Produce json
// @Param request body UpdateAuthorRequest true "Request body"
// @Param id query string true "Author identification"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /author/update [put]
func UpdateAuthorHandler(ctx *gin.Context) {

	authorId := ctx.Query("id")

	if authorId == "" {
		handler.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}


	request := UpdateAuthorRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		handler.Logger.ErrorF("validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	author := schemas.Author{}

	err = handler.Db.First(&author, authorId).Error

	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, "author not found")
		return 
	}

	if request.Age >= 0 {
		author.Age = request.Age
	}

	if request.Name != "" {
		author.Name = request.Name
	}

	err = handler.Db.Save(&author).Error

	if err != nil {
		handler.Logger.ErrorF("error updating author: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating author with id: %s", authorId))
		return
	}

	handler.SendSuccess(ctx, "update-author", author)





}