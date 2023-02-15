package server

func (application *Application) setupHandlers() {
	fiber := application.fiber
	fiber.Get("/status", statusHandler)
	fiber.Get("/v1/usage-statistics", application.getUsageStatisticsHandler())
}
