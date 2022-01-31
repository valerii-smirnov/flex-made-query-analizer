package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/services/dto"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/transport/definitions"
	"github.com/valerii-smirnov/flex-made-query-analizer/pkg/rest"
	"net/http"
	"strconv"
)

const (
	queryTypeParam = "type"
	sortingParam   = "sorting"
	pageParam      = "page"
	perPageParam   = "per-page"

	defaultPage    = "1"
	defaultPerPage = "20"
	defaultSorting = "first-slow"
)

// NewStatistic constructor for Statistic transport layer.
func NewStatistic(validator Validator, service StatisticService) *Statistic {
	return &Statistic{
		validator: validator,
		service:   service,
	}
}

// Statistic implements transport layer for Fiber application.
type Statistic struct {
	validator Validator
	service   StatisticService
}

// GetQueriesStatistic represents transport layer for.
// [GET] /database/queries
func (t Statistic) GetQueriesStatistic(ctx *fiber.Ctx) error {
	sPage := ctx.Query(pageParam, defaultPage)
	page, err := strconv.Atoi(sPage)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(rest.NewBadRequestError("error parsing page parameter from request", err))
	}

	sPerPage := ctx.Query(perPageParam, defaultPerPage)
	perPage, err := strconv.Atoi(sPerPage)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(rest.NewBadRequestError("error parsing per-page parameter from request", err))
	}

	req := definitions.GetQueriesRequest{
		QueryType: ctx.Query(queryTypeParam),
		Sorting:   ctx.Query(sortingParam, defaultSorting),
		Page:      page,
		PerPage:   perPage,
	}

	if err := t.validator.Struct(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(rest.NewBadRequestError("request validation error", err))
	}

	gc := dto.GetQueriesReq{
		QueryType: req.QueryType,
		Sorting:   req.Sorting,
		Page:      req.Page,
		PerPage:   req.PerPage,
	}

	data, err := t.service.GetStatistic(ctx.Context(), gc)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(rest.NewInternalServerError("error getting database queries statistic", err))
	}

	resp := make(definitions.GetQueriesResponse, 0, len(data))
	for _, row := range data {
		resp = append(resp, &definitions.QueryRow{
			QueryID:           row.QueryID,
			Query:             row.Query,
			MaxExecutionTime:  row.MaxExecutionTime,
			MeanExecutionTime: row.MeanExecutionTime,
		})
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}
