package repository

import (
	"cellular-data-tracker/model"
	"context"
	"database/sql"
	"fmt"
	"time"
)

const findQuery = "SELECT * FROM phones WHERE number=$1"

type postgresPhoneRepository struct {
	readTimeout time.Duration
	dB          *sql.DB
}

func NewPhoneRepository(connectionTimeout time.Duration, db *sql.DB) PhoneRepository {
	return &postgresPhoneRepository{readTimeout: connectionTimeout, dB: db}
}

func (repository postgresPhoneRepository) Find(number string) (*model.Phone, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repository.readTimeout)
	defer cancel()
	queryContext, err := repository.dB.QueryContext(ctx, findQuery, number)
	if err != nil {
		return nil, err
	}
	if queryContext.Next() {
		phone := &model.Phone{}
		err := queryContext.Scan(&phone.Id, &phone.Number)
		if err != nil {
			return nil, err
		} else {
			return phone, nil
		}
	} else {
		return nil, fmt.Errorf("unable to find phone with number [%s]", number)
	}
}
