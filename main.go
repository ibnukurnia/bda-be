package main

import (
	"fmt"
	"msbda/pkg/config"
	_ "msbda/pkg/env"
	"msbda/src/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	key := os.Getenv("APP_KEY")

	if key == "" {
		fmt.Println("Please Set APP KEY")
		os.Exit(1)
	}

	e := gin.Default()
	config.Cors(e, func(c *cors.Config) {
		c.AllowMethods = []string{"GET", "POST"}
		c.AllowOrigins = []string{"http://localhost:3000"}
		c.AllowHeaders = []string{"authorization"}
	})

	routes.Api(e)

	e.Run(":8000")
}
