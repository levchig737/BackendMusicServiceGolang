package controllers

import (
	"BackendMusicServiceGolang/models"
	"BackendMusicServiceGolang/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	songService *services.SongService
}

func NewSongController(service *services.SongService) *SongController {
	return &SongController{service}
}

// GetAllSongs retrieves all songs with filtering and pagination.
// @Summary Get all songs
// @Description Retrieve a list of songs with optional filters
// @Produce json
// @Param group query string false "Filter by group"
// @Param title query string false "Filter by title"
// @Param text query string false "Filter by text"
// @Param limit query int false "Limit of results"
// @Param page query int false "Page number"
// @Success 200 {array} models.Song
// @Failure 500 {object} map[string]interface{}
// @Router /songs [get]
func (ctrl *SongController) GetAllSongs(c *gin.Context) {
	// Фильтры
	filter := map[string]interface{}{}
	if group := c.Query("group"); group != "" {
		filter["group"] = group
	}
	if title := c.Query("title"); title != "" {
		filter["title"] = title
	}
	if text := c.Query("text"); text != "" {
		filter["text"] = text
	}

	// Пагинация
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	songs, err := ctrl.songService.GetAllSongs(filter, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении песен"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// AddSong adds a new song.
// @Summary Add a new song
// @Description Add a new song with the specified group and title
// @Accept json
// @Produce json
// @Param song body models.SongInput true "Song info"
// @Success 200 {object} models.Song
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /songs [post]
func (ctrl *SongController) AddSong(c *gin.Context) {
	var songInput models.SongInput

	if err := c.ShouldBindJSON(&songInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song := models.Song{
		Group: songInput.Group,
		Title: songInput.Song,
	}

	if err := ctrl.songService.Repo.AddSong(song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить песню"})
		return
	}

	c.JSON(http.StatusCreated, song)
}

// DeleteSong deletes a song by ID.
// @Summary Delete a song
// @Description Delete a song by its ID
// @Param id path int true "Song ID"
// @Success 204
// @Failure 500 {object} map[string]interface{}
// @Router /songs/{id} [delete]
func (ctrl *SongController) DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := ctrl.songService.Repo.DeleteSong(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдена"})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateSong updates an existing song.
// @Summary Update a song
// @Description Update a song's details
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body models.Song true "Updated song info"
// @Success 200 {object} models.Song
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /songs/{id} [put]
func (ctrl *SongController) UpdateSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var songInput models.Song

	if err := c.ShouldBindJSON(&songInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songInput.ID = uint(id)
	if err := ctrl.songService.Repo.UpdateSong(songInput); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдена"})
		return
	}

	c.JSON(http.StatusOK, songInput)
}

// GetSongTextByVerses retrieves the song text by verses.
// @Summary Get song text by verses
// @Description Get the text of a song with pagination
// @Produce json
// @Param id path int true "Song ID"
// @Param versesPerPage query int false "Verses per page"
// @Param page query int false "Page number"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /songs/{id}/text [get]
func (ctrl *SongController) GetSongTextByVerses(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	versesPerPage, _ := strconv.Atoi(c.DefaultQuery("versesPerPage", "4"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	text, err := ctrl.songService.GetSongTextByVerses(uint(id), versesPerPage, page)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ошибка при получении текста песни"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"text": text})
}
