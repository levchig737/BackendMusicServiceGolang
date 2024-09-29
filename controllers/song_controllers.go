package controllers

import (
	"BackendMusicServiceGolang/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	songService *services.SongService
}

func NewSongController(service *services.SongService) *SongController {
	return &SongController{service}
}

func (ctrl *SongController) AddSong(c *gin.Context) {
	var songInput struct {
		Group string `json:"group" binding:"required"`
		Song  string `json:"song" binding:"required"`
	}

	if err := c.ShouldBindJSON(&songInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err := ctrl.songService.FetchSongDetails(songInput.Group, songInput.Song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные о песне"})
		return
	}
	err = ctrl.songService.apiUrl
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить песню"})
		return
	}

	c.JSON(http.StatusOK, song)
}
