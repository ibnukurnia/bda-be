package anomaly_predictions

import (
	"msbda/pkg/app"
	"msbda/src/dtos/responses"

	"github.com/gookit/goutil/dump"
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

	ctx := app.GetCtx()

	dump.P(ctx.GetQueryIntOrDefault("page", 0))

	return anomalies, nil
}
