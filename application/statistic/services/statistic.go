package services

import (
	"context"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/repositories/models"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/services/dto"
)

// Statistic service implementation.
type Statistic struct {
	repository StatisticRepository
}

// NewStatistic constructor produces Statistic.
func NewStatistic(repository StatisticRepository) *Statistic {
	return &Statistic{repository: repository}
}

// GetStatistic receives dto.GetQueriesReq and calls repository to get statistic. Returns dto.QueryStatisticCollection.
func (s Statistic) GetStatistic(ctx context.Context, req dto.GetQueriesReq) (dto.QueryStatisticCollection, error) {

	filter := models.GetStatisticFilter{
		QueryType: req.QueryType,
		Sorting:   req.Sorting,
		Page:      req.Page,
		PerPage:   req.PerPage,
	}

	statisticCollection, err := s.repository.GetStatistic(ctx, filter)
	if err != nil {
		return nil, err
	}

	collection := make(dto.QueryStatisticCollection, 0, len(statisticCollection))
	for _, item := range statisticCollection {
		collection = append(collection, &dto.QueryStatistic{
			QueryID:           item.QueryID,
			Query:             item.Query,
			MaxExecutionTime:  item.MaxExecutionTime,
			MeanExecutionTime: item.MeanExecutionTime,
		})
	}

	return collection, nil
}
