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
	Repo   *repositories.SongRepository
	ApiUrl string
}

func NewSongService(repo *repositories.SongRepository, config config.Config) *SongService {
	return &SongService{Repo: repo, ApiUrl: config.ApiUrl}
}

// Получение всех песен с фильтрацией и пагинацией
func (service *SongService) GetAllSongs(filter map[string]interface{}, limit, offset int) ([]models.Song, error) {
	return service.Repo.GetAllSongs(filter, limit, offset)
}

// Получение текста песни по куплетам с пагинацией
func (service *SongService) GetSongTextByVerses(id uint, versesPerPage int, page int) (string, error) {
	return service.Repo.GetSongTextByVerses(id, versesPerPage, page)
}

func (service *SongService) FetchSongDetails(group string, song string) (*models.Song, error) {
	url := fmt.Sprintf("%s?group=%s&song=%s", service.ApiUrl, group, song)
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
		Text        string    `json:"text"`
		Link        string    `json:"link"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, err
	}

	return &models.Song{
		Group:       group,
		Title:       song,
		ReleaseDate: songDetail.ReleaseDate,
		Text:        songDetail.Text,
		Link:        songDetail.Link,
	}, nil
}
