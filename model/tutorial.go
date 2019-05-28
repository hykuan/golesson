package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Tutorial struct {
	ID        int `gorm:"primary_key"`
	Title     string
	AuthorID  int `json:"author_id"`
	Comments  []Comment `gorm:"foreignkey:TutorialID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	db, err := gorm.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&Tutorial{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Author{})
}
