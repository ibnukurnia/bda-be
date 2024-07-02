package anomaly_predictions

import (
	"msbda/pkg/app"
	"msbda/src/dtos/responses"
	"msbda/src/services/anomaly_predictions/data"
)

func HistoricalAnomaly() ([]responses.HistoricalAnomaly, error) {
	ctx := app.GetCtx()
	res := data.HistoricalAnomalies

	page := ctx.GetQueryIntOrDefault("page", 1)
	limit := ctx.GetQueryIntOrDefault("limit", 10)

	if limit < len(data.HistoricalAnomalies) {
		offset := func(l, p int) int {
			if p > 1 {
				return l
			}

			return 0
		}(limit, page)

		tail := func(d, l, p int) int {
			if l*p < d {
				return l * p
			}

			return d
		}(len(data.HistoricalAnomalies), limit, page)

		res = res[offset:tail]
	}

	return res, nil
}
