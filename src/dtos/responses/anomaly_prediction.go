package responses

type HistoricalAnomaly struct {
	Id           int    `json:"id"`
	ImpactedDate string `json:"impacted_date"`
	Severity     string `json:"severity"`
	Service      string `json:"service"`
	Description  string `json:"description"`
	TotalAlerts  int    `json:"total_alerts"`
	Status       string `json:"status"`
	Assignees    string `json:"assignees"`
}

type MostRecentAnomaly struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

type SeverityRatio struct {
	Severity   string  `json:"severity"`
	Percentage float32 `json:"percentage"`
}
