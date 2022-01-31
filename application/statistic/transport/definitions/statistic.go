package definitions

type GetQueriesRequest struct {
	QueryType string `validate:"omitempty,oneof=select insert update delete"`
	Sorting   string `validate:"oneof=first-slow first-fast,required"`
	Page      int    `validate:"gt=0,required"`
	PerPage   int    `validate:"gt=0,required"`
}

type GetQueriesResponse []*QueryRow

type QueryRow struct {
	QueryID           int64   `json:"query_id"`
	Query             string  `json:"query"`
	MaxExecutionTime  float64 `json:"max_execution_time"`
	MeanExecutionTime float64 `json:"mean_execution_time"`
}
