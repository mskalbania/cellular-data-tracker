package client

type CellularDataUsageStatisticsClient interface {
	GetUsageStatistics(phone string)
}
