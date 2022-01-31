package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/configuration"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/repositories"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/services"
	"github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/transport"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := logrus.New()

	cfg, err := configuration.New()
	if err != nil {
		logger.Fatal(logger.WithField("error", err))
	}

	db, err := gorm.Open(postgres.Open(cfg.DBDonfig.GetPostgresDsn()), &gorm.Config{})
	if err != nil {
		logger.Fatal(logger.WithField("error", err))
	}

	val := validator.New()
	statisticRepo := repositories.NewStatistic(db)
	statisticService := services.NewStatistic(statisticRepo)
	statisticTransport := transport.NewStatistic(val, statisticService)

	app := fiber.New()
	app.Get("/database/queries", statisticTransport.GetQueriesStatistic)

	if err := app.Listen(fmt.Sprintf(":%s", cfg.ApplicationPort)); err != nil {
		panic(err)
	}
}
