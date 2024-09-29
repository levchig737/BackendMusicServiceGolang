package models

import "time"

type Song struct {
	ID          uint      `gorm:"primaryKey"`
	Group       string    `json:"group"`
	Title       string    `json:"song"`
	ReleaseDate time.Time `json:"releaseDate"`
	Lyrics      string    `json:"lyrics"`
	Link        string    `json:"link"`
}
