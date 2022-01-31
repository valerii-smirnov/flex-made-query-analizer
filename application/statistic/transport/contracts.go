package transport

import (
	"context"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/services/dto"
)

//go:generate mockgen -destination=./contracts_mock_test.go -package=transport -source=./contracts.go

// Validator contract.
type Validator interface {
	Struct(s interface{}) error
}

// StatisticService contract.
type StatisticService interface {
	GetStatistic(ctx context.Context, req dto.GetQueriesReq) (dto.QueryStatisticCollection, error)
}
