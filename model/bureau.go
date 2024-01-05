package model

import "gorm.io/gorm"

type Bureau struct {
	gorm.Model
	UserID                int
	User                  User
	BureauReferenceNumber string
	UserReferenceNumber   string
	DiffNumM2M8           int
	CountPlLoan           int
	NewBureauVariable     string `gorm:"type:json"`
}
