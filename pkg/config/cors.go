package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type corsConfigFunc func(*cors.Config)

func Cors(e *gin.Engine, conf ...corsConfigFunc) {
	defaultConfig := cors.DefaultConfig()

	for _, c := range conf {
		c(&defaultConfig)
	}

	e.Use(cors.New(defaultConfig))
}
