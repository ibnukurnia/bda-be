package paginator

type recordPaginated[T any] struct {
	Rows       []T `json:"rows"`
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
	Page       int `json:"page"`
	// LastPage  int `json:"last_page"`
}

type RequestPagination struct {
	Page, Limit int
	Queries     map[string]string
}

func New[T any](data []T) *recordPaginated[T] {
	totalRows := len(data)

	return &recordPaginated[T]{
		Rows:      data,
		TotalRows: totalRows,
	}
}
