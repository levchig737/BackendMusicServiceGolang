package models

import "time"

// Song represents a song entity.
// @Description Song model
type Song struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Group       string    `json:"group"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

type SongInput struct {
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}
