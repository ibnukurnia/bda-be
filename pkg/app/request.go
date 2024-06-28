package app

import (
	ep "msbda/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RequestValidated interface {
	Messages() map[string]string
}

func validateJson(c *gin.Context, request RequestValidated) error {
	err := c.ShouldBindJSON(request)
	messages := request.Messages()

	if err != nil {
		vErrors := err.(validator.ValidationErrors)
		for _, e := range vErrors {
			if m, ok := messages[e.Field()+"."+e.Tag()]; ok {
				return &ep.SchemaError{
					Status:  400,
					Message: m,
				}
			} else {
				return &ep.SchemaError{
					Status:  400,
					Message: e.StructField() + e.Tag(),
				}
			}
		}
	}

	return nil
}
