package overviews

import (
	"math/rand"
	"msbda/src/dtos/responses"
)

func TeamOverview() (responses.TeamOverview, error) {

	return responses.TeamOverview{
		Solved:     rand.Intn(500),
		OnProgress: rand.Intn(100),
		TeamMember: rand.Intn(100),
		Overviews: []responses.TeamOverviewSituation{
			{
				Id:               1190,
				ServiceName:      "bridgtl-sms-service",
				ImpactedDuration: "12 min",
				OpenIssues:       rand.Intn(5),
				Contributor:      rand.Intn(5),
				AlertAttempt:     rand.Intn(10),
			},
			{
				Id:               1191,
				ServiceName:      "bridgtl-sms-service",
				ImpactedDuration: "19 sec",
				OpenIssues:       rand.Intn(5),
				Contributor:      rand.Intn(5),
				AlertAttempt:     rand.Intn(10),
			},
			{
				Id:               1192,
				ServiceName:      "bridgtl-sms-service",
				ImpactedDuration: "1 hr",
				OpenIssues:       rand.Intn(5),
				Contributor:      rand.Intn(5),
				AlertAttempt:     rand.Intn(10),
			},
		},
	}, nil
}
