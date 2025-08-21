package db

import (
	"github.com/x01000101x/ticket-project/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Ticket{})
}
