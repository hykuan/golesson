package model

import "time"

type Author struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Tutorials []Tutorial `gorm:"foreignkey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
