package repository

import (
	"cellular-data-tracker/model"
	"context"
	"database/sql"
	"time"
)

const insertQuery = "INSERT INTO usage_statistics(id,phoneId,timestamp,amount) VALUES($1,$2,$3,$4)"

type postgresUsageRecordRepository struct {
	readTimeout time.Duration
	dB          *sql.DB
}

func NewCellularDataStatisticsRepository(connectionTimeout time.Duration, db *sql.DB) UsageRecordRepository {
	return &postgresUsageRecordRepository{readTimeout: connectionTimeout, dB: db}
}

func (repository postgresUsageRecordRepository) Save(record *model.UsageRecord) error {
	ctx, cancel := context.WithTimeout(context.Background(), repository.readTimeout)
	defer cancel()
	_, err := repository.dB.ExecContext(ctx, insertQuery, record.Id, record.PhoneId, record.Timestamp, record.Amount)
	if err != nil {
		return err
	}
	return nil
}
