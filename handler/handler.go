package handler

import (
	"github.com/PedroBSanchez/gobooks.git/config"
	"gorm.io/gorm"
)


var (
	Logger *config.Logger
	Db *gorm.DB
)

func InitializeHandler() {
    Logger = config.GetLogger("handler")
	Db = config.GetSQLite()
}