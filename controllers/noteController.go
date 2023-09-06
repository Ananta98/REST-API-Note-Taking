package controllers

import (
	"net/http"
	"rest-api-note-taking/models"
	"rest-api-note-taking/utils/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NoteInput struct {
	Title   string `gorm:"size:255" json:"title"`
	Content string `gorm:"text" json:"content"`
}

func CreateNewNote(ctx *gin.Context) {
	input := models.Note{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	newNote := models.Note{
		UserID:  user_id,
		Title:   input.Title,
		Content: input.Title,
	}
	createdNote := models.Note{}
	if err := db.Create(&newNote).Last(&createdNote).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Create new note success", "data": createdNote})
}

func DeleteNote(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	var note models.Note
	if err := db.Where("id = ?", id).Find(&note).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Delete(&note).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete note success"})
}

func UpdateNote(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	var note models.Note
	if err := db.Where("id = ?", id).Find(&note).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := NoteInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedNoteResult := models.Note{}
	updateNote := models.Note{
		Title:   input.Title,
		Content: input.Content,
	}
	if err := db.Model(&note).Updates(&updateNote).Find(&updatedNoteResult).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Update note success", "data": updatedNoteResult})
}

func GetListNotes(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var notes []models.Note
	if err := db.Model("notes").Where("user_id = ?", user_id).Find(&notes).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success get all note list", "data": notes})
}

func GetDetailNote(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	var note models.Note
	if err := db.Model("notes").Where("id = ?", id).Find(&note).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success get all note list", "data": note})
}
