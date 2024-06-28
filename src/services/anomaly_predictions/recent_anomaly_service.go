package anomaly_predictions

import (
	"msbda/src/dtos/responses"
)

func MostRecentAnomaly() ([]responses.MostRecentAnomaly, error) {
	anomalies := []responses.MostRecentAnomaly{
		{
			Name:  "Pods",
			Total: 100,
		},
		{
			Name:  "VM",
			Total: 75,
		},
		{
			Name:  "Databases",
			Total: 50,
		},
		{
			Name:  "Services",
			Total: 25,
		},
	}

	return anomalies, nil
}
