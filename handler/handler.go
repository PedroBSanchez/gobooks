package handler

import (
	"github.com/PedroBSanchez/gobooks.git/config"
	"gorm.io/gorm"
)


var (
	Logger *config.Logger
	Db *gorm.DB
	CustomLayout string
)

func InitializeHandler() {
    Logger = config.GetLogger("handler")
	Db = config.GetSQLite()
	CustomLayout = "2006-01-02 15:04:05"
}