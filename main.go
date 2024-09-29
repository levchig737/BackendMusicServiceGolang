package main

import (
	"BackendMusicServiceGolang/config"
	"BackendMusicServiceGolang/controllers"
	"BackendMusicServiceGolang/repositories" // Добавь импорт для репозитория
	"BackendMusicServiceGolang/routes"
	"BackendMusicServiceGolang/services" // Добавь импорт для сервиса
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig() // Избегаем конфликта с названием пакета

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Создание экземпляра репозитория
	repo := repositories.NewSongRepository(db)

	// Создание сервиса песни
	songService := services.NewSongService(repo, cfg)

	// Настройка роутера
	router := gin.Default()

	// Инициализация маршрутов
	routes.SetupRoutes(router, controllers.NewSongController(songService))

	// Запуск сервера
	router.Run(":" + cfg.Port)
}
