package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func InitializeSQLite() (*gorm.DB, error) {
	
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()

	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate schemas
	//err = db.AutoMigrate(&schemas.Opening{})
	//if err != nil {
		//logger.ErrorF("sqlite automigration error; %v", err)
		//return nil, err
	//}

	// Return DB
	return db, nil

}