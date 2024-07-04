package middleware

import (
	"errors"
	"msbda/pkg/app"
	"msbda/pkg/e"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		appG := app.Context(c)

		if bearer == "" {
			appG.Response(401, nil).
				WithMessage("unautozrized").
				Send()

			c.Abort()
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
			appG.Response(401, nil).
				WithMessage(err.Error()).
				Send()

			c.Abort()
			return
		}

		_, OK := t.Claims.(jwt.MapClaims)
		if !OK {
			appG.Response(e.ERROR, nil).
				WithMessage("failed to parse token").
				Send()

			c.Abort()
			return
		}

		c.Next()
	}
}
