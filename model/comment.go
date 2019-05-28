package model

import (
	"time"
)

type Comment struct {
	ID         int `gorm:"primary_key"`
	TutorialID int
	AuthorID   int `json:"author_id"`
	Body       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}
