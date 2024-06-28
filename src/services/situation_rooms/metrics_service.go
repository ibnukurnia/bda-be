package situation_rooms

import "math/rand"

type Metric struct {
	Data map[string]int `json:"data"`
}

func GetMetricSituation() (*Metric, error) {
	return &Metric{
		Data: map[string]int{
			"Wed, 11 Jan 2024 12:03:56": rand.Intn(250),
			"Wed, 12 Jan 2024 12:03:56": rand.Intn(250),
			"Wed, 13 Jan 2024 12:03:56": rand.Intn(250),
		},
	}, nil
}
