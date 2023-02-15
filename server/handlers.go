package server

import (
	"cellular-data-tracker/server/api"
	"github.com/gofiber/fiber/v2"
)

func statusHandler(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&api.Status{Status: "UP"})
}

func (application *Application) getUsageStatisticsHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		application.cellularDataService.GetAllUsageStatistic()
		//serialization + logging + return
		ctx.Status(200)
		return nil
	}
}
