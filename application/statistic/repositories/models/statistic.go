package models

// GetStatisticFilter request parameter represents filter.
type GetStatisticFilter struct {
	QueryType string
	Sorting   string
	Page      int
	PerPage   int
}

// GetStatisticResultRow model for mapping database result.
type GetStatisticResultRow struct {
	QueryID           int64   `gorm:"column:queryid"`
	Query             string  `gorm:"column:query"`
	MaxExecutionTime  float64 `gorm:"column:max_exec_time"`
	MeanExecutionTime float64 `gorm:"column:mean_exec_time"`
}

// TableName provides mode table name.
func (g GetStatisticResultRow) TableName() string {
	return "pg_stat_statements"
}

// GetStatisticResultCollection collection of GetStatisticResultRow
type GetStatisticResultCollection []*GetStatisticResultRow
