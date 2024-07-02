package handlers

import (
	"msbda/pkg/app"
	"msbda/pkg/helpers/paginator"
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

func HistoricalAnomalies(c *gin.Context) {
	appG := app.Context(c)

	appG.WithQueries("page", "limit")

	res, err := anomaly_predictions.HistoricalAnomaly()
	if !appG.ParseError(err) {
		return
	}

	r := paginator.New(res)

	appG.Response(200, r).
		Send()
}
