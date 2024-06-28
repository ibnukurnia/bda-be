package handlers

import (
	"msbda/pkg/app"
	"msbda/src/services/anomaly_predictions"

	"github.com/gin-gonic/gin"
)

func RecentAnomaly(c *gin.Context) {
	appG := app.Context(c)

	recentAnomalies, err := anomaly_predictions.MostRecentAnomaly()
	if !appG.ParseError(err) {
		return
	}

	appG.Response(200, recentAnomalies).
		Send()
}

func SeverityRatio(c *gin.Context) {
	appG := app.Context(c)

	res, err := anomaly_predictions.SeverityRatio()
	if !appG.ParseError(err) {
		return
	}

	appG.Response(200, res).
		Send()
}
