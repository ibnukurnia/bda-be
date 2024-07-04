package handlers

import (
	"math/rand"
	"msbda/pkg/app"
	"msbda/src/dtos/responses"
	"msbda/src/services/overviews"

	"github.com/gin-gonic/gin"
)

func Insights(c *gin.Context) {
	appG := app.Context(c)

	res := responses.Insight{
		EventTriggered:   rand.Intn(3000),
		Alerts:           rand.Intn(500),
		OnGoingSituation: rand.Intn(100),
		AvgTimeSolved:    rand.Intn(30),
		CurrentSituation: []responses.SummarySituation{
			{
				Id:          1,
				CreatedDate: "11/01/2024 15:04:22 P.M",
				Severity:    "minor",
				Status:      "solved",
				Assignees: []responses.SituationRoomAssignee{
					{
						Email: "davin@bri.co.id",
					},
					{
						Email: "davin1@bri.co.id",
					},
				},
				Service:     "bridgtl-rsm-notifications",
				Description: "A clash occurred while attempting to save n..",
				TotalAlerts: 7,
			},
			{
				Id:          1,
				CreatedDate: "11/01/2024 15:04:22 P.M",
				Severity:    "minor",
				Status:      "open",
				Assignees: []responses.SituationRoomAssignee{
					{
						Email: "davin@bri.co.id",
					},
				},
				Service:     "bridgtl-rsm-notifications",
				Description: "A clash occurred while attempting to save n..",
				TotalAlerts: 7,
			},
			{
				Id:          1,
				CreatedDate: "11/01/2024 15:04:22 P.M",
				Severity:    "minor",
				Status:      "in progress",
				Assignees:   []responses.SituationRoomAssignee{},
				Service:     "bridgtl-rsm-notifications",
				Description: "A clash occurred while attempting to save n..",
				TotalAlerts: 7,
			},
			{
				Id:          1,
				CreatedDate: "11/01/2024 15:04:22 P.M",
				Severity:    "minor",
				Status:      "solved",
				Assignees: []responses.SituationRoomAssignee{
					{
						Email: "davin@bri.co.id",
					},
					{
						Email: "davin1@bri.co.id",
					},
					{
						Email: "davin2@bri.co.id",
					},
				},
				Service:     "bridgtl-rsm-notifications",
				Description: "A clash occurred while attempting to save n..",
				TotalAlerts: 7,
			},
			{
				Id:          1,
				CreatedDate: "11/01/2024 15:04:22 P.M",
				Severity:    "minor",
				Status:      "solved",
				Assignees: []responses.SituationRoomAssignee{
					{
						Email: "davin@bri.co.id",
					},
					{
						Email: "davin1@bri.co.id",
					},
					{
						Email: "davin2@bri.co.id",
					},
				},
				Service:     "bridgtl-rsm-notifications",
				Description: "A clash occurred while attempting to save n..",
				TotalAlerts: 7,
			},
		},
	}

	appG.Response(200, res).
		Send()
}

func ServiceOverview(c *gin.Context) {
	appG := app.Context(c)

	res, err := overviews.ServiceOverview()

	if !appG.ParseError(err) {
		return
	}

	appG.Response(200, res).
		Send()
}

func TeamOverview(c *gin.Context) {
	appG := app.Context(c)

	res, err := overviews.TeamOverview()

	if !appG.ParseError(err) {
		return
	}

	appG.Response(200, res).
		Send()
}
