package main

// Base Model's definition
import (
	"time"

	"github.com/jinzhu/gorm"
)

// Model to be imported in all the individual models.
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User Model.
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"type:varchar(100);unique_index"`
}
