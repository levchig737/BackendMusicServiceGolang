package repositories

import (
	"BackendMusicServiceGolang/models"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db}
}

// Получение всех песен с фильтрацией и пагинацией
func (repo *SongRepository) GetAllSongs(filter map[string]interface{}, limit int, offset int) ([]models.Song, error) {
	var songs []models.Song
	query := repo.db
	// Применение фильтров
	for key, value := range filter {
		query = query.Where("\""+key+"\""+" = ?", value)
	}

	// Пагинация
	err := query.Limit(limit).Offset(offset).Find(&songs).Error
	return songs, err
}

// Добавление новой песни
func (repo *SongRepository) AddSong(song models.Song) error {
	return repo.db.Create(&song).Error
}

// Обновление данных песни
func (repo *SongRepository) UpdateSong(song models.Song) error {
	return repo.db.Save(&song).Error
}

// Удаление песни по ID
func (repo *SongRepository) DeleteSong(id uint) error {
	return repo.db.Delete(&models.Song{}, id).Error
}

// Получение текста песни по куплетам/припевам
func (repo *SongRepository) GetSongTextByVerses(id uint, versesPerPage int, page int) (string, error) {
	var song models.Song
	if err := repo.db.First(&song, id).Error; err != nil {
		return "", err
	}

	text := song.Text
	verses := splitByVerses(text)
	start := (page - 1) * versesPerPage
	end := start + versesPerPage

	if start >= len(verses) {
		return "", nil
	}

	if end > len(verses) {
		end = len(verses)
	}

	return joinTextVerses(verses[start:end]), nil
}

// Функция для разбивки текста на куплеты/припевы по меткам
func splitByVerses(text string) []string {
	// Регулярное выражение для поиска маркеров куплетов и припевов
	re := regexp.MustCompile(`(?i)\[(припев|куплет)\]`)
	parts := re.Split(text, -1)

	// Извлекаем маркеры и соответствующие части текста
	verses := []string{}
	matches := re.FindAllString(text, -1)

	for i, part := range parts {
		part = strings.TrimSpace(part) // Убираем лишние пробелы
		if part != "" {
			if i < len(matches) {
				verses = append(verses, matches[i]+"\n"+part)
			} else {
				verses = append(verses, part)
			}
		}
	}

	return verses
}

// Функция для склейки куплетов/припевов
func joinTextVerses(verses []string) string {
	return strings.Join(verses, "\n\n") // Склеиваем с двойным переносом строк для разделения куплетов
}

// // Функция для разбивки текста на строки
// func splitText(text string) []string {
// 	return strings.Split(text, "\n")
// }

// // Функция для склейки строк
// func joinText(lines []string) string {
// 	return strings.Join(lines, "\n")
// }
