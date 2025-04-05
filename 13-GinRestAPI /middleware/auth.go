package middleware

import (
	"fmt"
	"net/http"

	"example.com/apis/utils"
	"github.com/gin-gonic/gin"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	fmt.Println("token===", token)

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is required"})
		return
	}
	userId, err := utils.VerifyToken(token)
	fmt.Println("tokenData===", userId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
