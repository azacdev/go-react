package database

import "github.com/azacdev/go-react/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Entry{})
}
