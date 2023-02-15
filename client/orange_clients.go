package client

type orangeCellularDataUsageStatisticsClient struct {
	//http client
}

func NewCellularDataUsageStatisticsClient() CellularDataUsageStatisticsClient {
	return &orangeCellularDataUsageStatisticsClient{}
}

func (client *orangeCellularDataUsageStatisticsClient) GetUsageStatistics(phone string) {
}
