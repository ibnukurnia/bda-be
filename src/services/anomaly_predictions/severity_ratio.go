package anomaly_predictions

import "msbda/src/dtos/responses"

func SeverityRatio() ([]responses.SeverityRatio, error) {
	severities := []responses.SeverityRatio{
		{
			Severity:   "Minor",
			Percentage: 40,
		},
		{
			Severity:   "Mayor",
			Percentage: 25,
		},
		{
			Severity:   "No Severity",
			Percentage: 35,
		},
	}

	return severities, nil
}
