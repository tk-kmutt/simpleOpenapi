package repository

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type AmazonItems struct {
	Name      string
	Maker     string
	Price     int64
	Comment   *string
	Url       string
	Asin      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDelete  soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
