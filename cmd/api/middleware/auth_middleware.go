package middleware

import (
	"douyin/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, ok := ctx.GetQuery("token")
		fmt.Println(">>>>>>>>>>>>>>>>>>>", tokenString)

		if tokenString == "" || !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"StatusCode": 401, "StatusMsg": "Insufficient permissions"})
			ctx.Abort()
			return
		}

		token, claims, err := middleware.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"StatusCode": 401, "StatusMsg": "Insufficient permissions"})
			ctx.Abort()
			return
		}

		userID := claims.UserID
		fmt.Println(">>>>>>>>>>>>>>>>>>>", userID)
		// TODO: search user in db
		/*if false {
			ctx.JSON(http.StatusUnauthorized, gin.H{"StatusCode": 401, "StatusMsg": "Insufficient permissions"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)*/
		ctx.Next()
	}
}
