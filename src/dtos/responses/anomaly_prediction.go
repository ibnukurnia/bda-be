package responses

type MostRecentAnomaly struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

type SeverityRatio struct {
	Severity   string  `json:"severity"`
	Percentage float32 `json:"percentage"`
}
