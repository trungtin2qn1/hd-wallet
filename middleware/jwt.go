package middleware

import (
	"fmt"
	"hd-wallet/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//VerifyJWTToken ...
func VerifyJWTToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	rawToken := string(token[len("Tin "):])
	userID, err := jwt.VerificationToken(rawToken)
	if err != nil {
		//go utils.LogErrToFile(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	fmt.Println(userID)
	c.Set("user_id", userID)
	c.Next()
}
