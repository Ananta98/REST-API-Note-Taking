package routes

import (
	"rest-api-note-taking/controllers"
	"rest-api-note-taking/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	auth := router.Group("auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)

	note := router.Group("note")
	note.Use(middlewares.Middlewares())
	note.GET("/", controllers.GetListNotes)
	note.POST("/", controllers.CreateNewNote)
	note.GET("/:id", controllers.GetDetailNote)
	note.PATCH("/:id", controllers.UpdateNote)
	note.DELETE("/:id", controllers.DeleteNote)
	return router
}
