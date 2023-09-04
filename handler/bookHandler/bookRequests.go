package bookhandler

import (
	"fmt"
	"time"

	"github.com/PedroBSanchez/gobooks.git/handler"
)


func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateBookRequest struct {
	Title string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
	Pages int64 `json:"pages"`
    AuthorID int64 `json:"authorID"`
}


func (r *CreateBookRequest) Validate() error{

	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}

	if r.Pages <= 0 {
		return errParamIsRequired("pages", "int64")
	}

	now := time.Now()
	timeDate, err := time.Parse(handler.CustomLayout, r.ReleaseDate)
	if  err != nil || now.Before(timeDate){
		return fmt.Errorf("Invalid date")
	}

	if r.AuthorID <= 0 {
		return errParamIsRequired("authorID", "int64")
	}

	return nil
}


type UpdateBookRequest struct {
	Title string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
	Pages int64 `json:"pages"`
    AuthorID int64 `json:"authorID"`
}


func (r *UpdateBookRequest) Validate() error {
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}

	if r.Pages <= 0 {
		return errParamIsRequired("pages", "int64")
	}

	now := time.Now()
	timeDate, err := time.Parse(handler.CustomLayout, r.ReleaseDate)
	if  err != nil || now.Before(timeDate){
		return fmt.Errorf("Invalid date")
	}

	if r.AuthorID <= 0 {
		return errParamIsRequired("authorID", "int64")
	}

	return nil
}