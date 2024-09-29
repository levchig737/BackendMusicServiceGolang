package services

import (
	"BackendMusicServiceGolang/config"
	"BackendMusicServiceGolang/models"
	"BackendMusicServiceGolang/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SongService struct {
	repo   *repositories.SongRepository
	apiUrl string
}

func NewSongService(repo *repositories.SongRepository, config config.Config) *SongService {
	return &SongService{repo, config.ApiUrl}
}

func (service *SongService) FetchSongDetails(group string, song string) (*models.Song, error) {
	url := fmt.Sprintf("%s?group=%s&song=%s", service.apiUrl, group, song)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("не удалось получить данные о песне")
	}

	var songDetail struct {
		ReleaseDate time.Time `json:"releaseDate"`
		Lyrics      string    `json:"text"`
		Link        string    `json:"link"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, err
	}

	return &models.Song{
		Group:       group,
		Title:       song,
		ReleaseDate: songDetail.ReleaseDate,
		Lyrics:      songDetail.Lyrics,
		Link:        songDetail.Link,
	}, nil
}
