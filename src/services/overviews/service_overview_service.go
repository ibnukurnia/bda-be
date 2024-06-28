package overviews

import (
	"msbda/src/dtos/responses"
)

func ServiceOverview() ([]responses.ServiceOverview, error) {

	return []responses.ServiceOverview{
		{
			SituationIds:     []int{1190, 1192, 1028},
			ServiceName:      "bridgtl-sms-service",
			OpenIssues:       3,
			Contributor:      30,
			CurrentCondition: "Running",
			LastImpacted:     "12 min",
		},
		{
			SituationIds:     []int{2029},
			ServiceName:      "bridgtl-trx-service",
			OpenIssues:       1,
			Contributor:      30,
			CurrentCondition: "Crash",
			LastImpacted:     "23 min",
		},
		{
			SituationIds:     []int{2028},
			ServiceName:      "load-balancer",
			OpenIssues:       1,
			Contributor:      2,
			CurrentCondition: "Warning",
			LastImpacted:     "45 sec",
		},
	}, nil
}
