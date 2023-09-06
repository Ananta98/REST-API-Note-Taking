package middlewares

import (
	"net/http"
	"rest-api-note-taking/utils/token"

	"github.com/gin-gonic/gin"
)

func Middlewares() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := token.ValidToken(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
