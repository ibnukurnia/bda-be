package handlers

import (
	"msbda/pkg/app"
	"msbda/src/dtos/responses"
	"msbda/src/services/situation_rooms"

	"github.com/gin-gonic/gin"
)

func OnGoingSituation(c *gin.Context) {
	appG := app.Context(c)

	situations := []responses.SummarySituation{
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
			Status:      "open",
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
			Status:      "in progress",
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
	}

	appG.Response(200, situations).Send()
}

func SituationAlerts(c *gin.Context) {
	appG := app.Context(c)

	id, s := appG.GetParamInt("id")
	if !s {
		return
	}

	res, err := situation_rooms.AlertService(id)

	if !appG.ParseError(err) {
		return
	}

	appG.Response(200, res).
		Send()
}

func Metrics(c *gin.Context) {
	appG := app.Context(c)

	res, err := situation_rooms.GetMetricSituation()

	if !appG.ParseError(err) {
		return
	}

	appG.Response(200, gin.H{
		"metric": res,
	}).
		Send()
}

func Topology(c *gin.Context) {
	appG := app.Context(c)

	id, valid := appG.GetParamInt("id")
	if !valid {
		return
	}

	res, err := situation_rooms.TopologySituation(id)

	if !appG.ParseError(err) {
		return
	}

	appG.Response(200, res).
		Send()
}
