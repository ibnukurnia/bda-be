package app

import (
	"msbda/pkg/e"

	"github.com/gin-gonic/gin"
)

type response struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    any          `json:"data"`
	valid   bool         `json:"-"`
	ctx     *gin.Context `json:"-"`
}

func (r *response) WithMessage(message string) *response {
	r.Message = message

	return r
}

func (r *response) Send() {
	r.valid = r.Status >= 200 || r.Status <= 299

	if r.Message == "" {
		r.Message = e.GetMsg(r.Status)
	}

	r.ctx.JSON(r.Status, r)
}
