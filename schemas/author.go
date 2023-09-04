package schemas

import (
	"time"

	"gorm.io/gorm"
)


type Author struct {
	gorm.Model
	Name string
	Age int64
	AmountBooks int64
	Books []Book `gorm:"foreignkey:AuthorID;constraint:OnDelete:CASCADE;"`
}


type AuthorResponse struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:deletedAt, omitempty`
	Name string `json:"name"`
	Age int64  `json:"age"`
	AmountBooks int64 `json:"amountBooks"`
}