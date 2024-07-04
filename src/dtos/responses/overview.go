package responses

type Insight struct {
	EventTriggered   int                `json:"event_triggered"`
	Alerts           int                `json:"total_alerts"`
	OnGoingSituation int                `json:"on_going_situation"`
	AvgTimeSolved    int                `json:"avg_time_solved"`
	CurrentSituation []SummarySituation `json:"current_situations"`
}

type TeamOverview struct {
	Solved     int                     `json:"solved"`
	OnProgress int                     `json:"on_progress"`
	TeamMember int                     `json:"team_member"`
	Overviews  []TeamOverviewSituation `json:"overviews"`
}

type TeamOverviewSituation struct {
	Id               int    `json:"id"`
	ServiceName      string `json:"service_name"`
	ImpactedDuration string `json:"impacted_duration"`
	OpenIssues       int    `json:"open_issues"`
	Contributor      int    `json:"contributor"`
	AlertAttempt     int    `json:"alert_attempt"`
}

type ServiceOverview struct {
	SituationIds     []int  `json:"situation_ids"`
	ServiceName      string `json:"service_name"`
	OpenIssues       int    `json:"open_issues"`
	Contributor      int    `json:"contributor"`
	CurrentCondition string `json:"current_condition"`
	LastImpacted     string `json:"last_impacted"`
}

type MetricOverview struct {
	ServiceName string `json:"service_name"`
}
