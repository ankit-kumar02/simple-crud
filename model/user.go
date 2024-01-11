package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Email    string
	Age      int
	Birthday *time.Time
	Address  string
}
