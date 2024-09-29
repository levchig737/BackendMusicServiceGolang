package repositories

import (
	"BackendMusicServiceGolang/models"

	"gorm.io/gorm"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db}
}

func (repo *SongRepository) GetAllSongs() ([]models.Song, error) {
	var songs []models.Song
	err := repo.db.Find(&songs).Error
	return songs, err
}

func (repo *SongRepository) AddSong(song models.Song) error {
	return repo.db.Create(&song).Error
}

func (repo *SongRepository) UpdateSong(song models.Song) error {
	return repo.db.Save(&song).Error
}

func (repo *SongRepository) DeleteSong(id uint) error {
	return repo.db.Delete(&models.Song{}, id).Error
}
