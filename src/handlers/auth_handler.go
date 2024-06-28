package handlers

import (
	"msbda/pkg/app"
	"msbda/src/dtos/requests"
	"msbda/src/services/auth"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	appG := app.Context(c)

	var req requests.LoginRequest

	if !appG.BindJson(&req) {
		return
	}

	res, err := auth.Login(req.Pernr, req.Password)

	if !appG.ParseError(err) {
		return
	}

	appG.Response(200, res).
		WithMessage("Loged in").
		Send()
}
