package token

import (
	"fmt"
	"rest-api-note-taking/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = utils.GetEnv("SECRET_KEY", "rahasiasekali")

func ExtractTokenString(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}
	header := ctx.Request.Header.Get("Authorization")
	splitted := strings.Split(header, " ")
	if len(splitted) == 2 {
		return splitted[1]
	}
	return ""
}

func CreateJwtToken(user_id uint) (string, error) {
	ttl, err := strconv.Atoi(utils.GetEnv("TOKEN_TIME_TO_TILE", "1"))
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["id"] = user_id
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(ttl)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}

func ValidToken(ctx *gin.Context) error {
	token := ExtractTokenString(ctx)
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractTokenID(ctx *gin.Context) {

}
