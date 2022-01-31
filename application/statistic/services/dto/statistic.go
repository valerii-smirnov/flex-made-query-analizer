package dto

// GetQueriesReq request data transfer object.
type GetQueriesReq struct {
	QueryType string
	Sorting   string
	Page      int
	PerPage   int
}

// QueryStatistic result dto.
type QueryStatistic struct {
	QueryID           int64
	Query             string
	MaxExecutionTime  float64
	MeanExecutionTime float64
}

// QueryStatisticCollection results collection.
type QueryStatisticCollection []*QueryStatistic
