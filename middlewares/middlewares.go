package middlewares

import "github.com/gin-gonic/gin"

func Middlewares() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
