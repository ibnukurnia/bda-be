package routes

import (
	"msbda/pkg/app"
	"msbda/src/handlers"
	"msbda/src/middleware"

	"github.com/gin-gonic/gin"
)

func Api(e *gin.Engine) {
	e.GET("/health", func(ctx *gin.Context) {
		appG := app.Context(ctx)

		appG.Response(200, gin.H{"message": "ok"}).Send()
	})

	api := e.Group("/api")
	{
		api.POST("/login", handlers.Login)

		api.Use(middleware.Auth())
		{
			situationRoomsEp := api.Group("situation-rooms")
			{
				situationRoomsEp.GET("/on-going", handlers.OnGoingSituation)
				detailEp := situationRoomsEp.Group("/:id")
				{
					detailEp.GET("/metrics", handlers.Metrics)
					detailEp.GET("/topology", handlers.Topology)
					detailEp.GET("/alerts", handlers.SituationAlerts)
				}
			}

			anomalyPredictionsEp := api.Group("anomaly-predictions")
			{
				anomalyPredictionsEp.GET("/most-recent", handlers.RecentAnomaly)
				anomalyPredictionsEp.GET("/severity-ratio", handlers.SeverityRatio)
			}

			overviewsEp := api.Group("overview")
			{
				overviewsEp.GET("/insights", handlers.Insights)
				overviewsEp.GET("/teams", handlers.TeamOverview)
				overviewsEp.GET("/services", handlers.ServiceOverview)
			}
		}
	}

}
