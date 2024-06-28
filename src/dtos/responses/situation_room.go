package responses

type SituationRoomAssignee struct {
	Email string `json:"email"`
}

type SummarySituation struct {
	Id          int                     `json:"id"`
	TotalAlerts int                     `json:"total_alerts"`
	CreatedDate string                  `json:"created_date"`
	Severity    string                  `json:"severity"`
	Status      string                  `json:"status"`
	Service     string                  `json:"service"`
	Description string                  `json:"description"`
	Assignees   []SituationRoomAssignee `json:"assignees"`
}

type SituationRoomAlert struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	AlertTime string `json:"alert_time"`
}

type TimelineSituation struct {
	Events    []TimelineEvent `json:"events"`
	Timestamp []string        `json:"time_stamps"`
}

type TimelineEvent struct {
	Timestamp TimelineTimestamp `json:"time_stamp"`
	Event     string            `json:"event"`
	Title     string            `json:"title"`
}

type TimelineTimestamp struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type TopologySituation struct {
	Nodes *NodeTopology `json:"nodes"`
}

type NodeTopology struct {
	Name   string          `json:"name"`
	Childs []*NodeTopology `json:"childs"`
}
