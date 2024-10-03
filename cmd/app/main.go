package main

import (
	"BackendMusicServiceGolang/config"
	"BackendMusicServiceGolang/controllers"
	"BackendMusicServiceGolang/models"
	"BackendMusicServiceGolang/repositories"
	"BackendMusicServiceGolang/routes"
	"BackendMusicServiceGolang/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "BackendMusicServiceGolang/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Music API
// @version 1.0
// @description API для управления музыкой
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
// @schemes http

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Миграция схемы базы данных
	err = db.AutoMigrate(
		&models.Song{},
	)
	if err != nil {
		log.Fatal("Ошибка миграции базы данных:", err)
	}

	// Создание экземпляра репозитория
	repo := repositories.NewSongRepository(db)

	// Создание сервиса песни
	songService := services.NewSongService(repo, cfg)

	// Настройка роутера
	router := gin.Default()

	// Инициализация маршрутов
	routes.SetupRoutes(router, controllers.NewSongController(songService))

	// Подключение Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	router.Run(":" + cfg.Port)
}
