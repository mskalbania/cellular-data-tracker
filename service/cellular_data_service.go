package service

import (
	"cellular-data-tracker/client"
	"cellular-data-tracker/model"
	"cellular-data-tracker/repository"
	"fmt"
)

type CellularDataService interface {
	GetAllUsageStatistic() ([]*model.Phone, error)
}

type cellularDataService struct {
	cellularDataClient    client.CellularDataUsageStatisticsClient
	usageRecordRepository repository.UsageRecordRepository
	phoneRepository       repository.PhoneRepository
}

func NewCellularDataService(
	cellularDataClient client.CellularDataUsageStatisticsClient,
	usageRecordRepository repository.UsageRecordRepository,
	phoneRepository repository.PhoneRepository,
) CellularDataService {
	return &cellularDataService{
		cellularDataClient:    cellularDataClient,
		usageRecordRepository: usageRecordRepository,
		phoneRepository:       phoneRepository,
	}
}

func (service *cellularDataService) GetAllUsageStatistic() ([]*model.Phone, error) {
	phone, err := service.phoneRepository.Find("505100477")
	if err != nil {
		return nil, err
	}
	fmt.Println(phone)
	return make([]*model.Phone, 1), nil
}
