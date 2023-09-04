package authorHandler

import (
	"fmt"
	"net/http"

	"github.com/PedroBSanchez/gobooks.git/handler"
	"github.com/PedroBSanchez/gobooks.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Delete author
// @Description Delete an author
// @Tags Author
// @Accept json
// @Produce json
// @Param id query string true "Author identification"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /author/delete [delete]
func DeleteAuthorHandler(ctx *gin.Context) {

	authorId := ctx.Query("id")

	if authorId == "" {
		handler.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}


	author := schemas.Author{}

	//Find author
	err := handler.Db.Preload("Books").First(&author, authorId).Error

	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("author with id: %s not found", authorId))
		return
	}

	// Excluir os livros relacionados
	for _, book := range author.Books {
		handler.Db.Delete(&book)
	}

	//Delete author
	err = handler.Db.Delete(&author).Error

	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting author with id: %s", authorId))
		return
	}

	handler.SendSuccess(ctx, "delete-author", author)

}