package models

import "time"

// Company table
type Company struct {
	// Date of create
	CreateDate *time.Time
	// company id of create
	CreateUID int
	// ID of company
	ID int `gorm:"primary_key:id"`
	// name of company
	Name string
	// Date of update
	UpdateDate *time.Time
	// company id of update
	UpdateUID int
}
