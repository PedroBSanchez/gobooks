package schemas

import (
	"time"

	"gorm.io/gorm"
)


type Book struct {
	gorm.Model
	Title string
	ReleaseDate time.Time
	Pages int64
	AuthorID uint
	Author Author

}


type BookResponse struct {
	ID uint `jsone:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:deletedAt, omitempty`
	Title string `json:"title"`
	ReleaseDate time.Time `json:"releaseDate"`
	Pages int64 `json:"pages"`
	AuthorId uint `json:"authorId"`
	Author AuthorResponse `json:"author"`
}