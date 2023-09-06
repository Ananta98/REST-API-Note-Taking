package models

import (
	"html"
	"strings"

	"rest-api-note-taking/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`

	// Relationship
	Notes []Note `json:"-"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func (u *User) SaveUser(db *gorm.DB) error {
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	if err := db.Save(&u).Error; err != nil {
		return err
	}
	return nil
}

func LoginValid(username, password string, db *gorm.DB) (string, error) {
	var foundUser User
	if err := db.Where("username = ?", username).Find(&foundUser).Error; err != nil {
		return "", err
	}
	err := VerifyPassword(password, foundUser.Password)
	if err != nil {
		return "", err
	}
	jwtToken, err := token.CreateJwtToken(foundUser.ID)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
