package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")

		if bearer == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unautorized",
			})

			return
		}

		token := strings.Split(bearer, " ")[1]

		t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, OK := t.Method.(*jwt.SigningMethodHMAC); !OK {
				return nil, errors.New("bad signed method received")
			}
			return []byte(os.Getenv("APP_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": err.Error(),
			})

			return
		}

		_, OK := t.Claims.(jwt.MapClaims)
		if !OK {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "unable to parse claims",
			})

			return
		}

		c.Next()
	}
}
