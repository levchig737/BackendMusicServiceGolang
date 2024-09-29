package routes

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, songController *controllers.SongController) {
	api := router.Group("/api")
	{
		api.POST("/songs", songController.AddSong)
		api.GET("/songs", songController.GetAllSongs)
		api.PUT("/songs/:id", songController.UpdateSong)
		api.DELETE("/songs/:id", songController.DeleteSong)
	}
}
