package services

import (
	"context"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/repositories/models"
)

//go:generate mockgen -destination=./contracts_mock_test.go -package=services -source=./contracts.go

// StatisticRepository contract.
type StatisticRepository interface {
	GetStatistic(ctx context.Context, filter models.GetStatisticFilter) (models.GetStatisticResultCollection, error)
}
