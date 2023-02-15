package repository

import "cellular-data-tracker/model"

type UsageRecordRepository interface {
	Save(record *model.UsageRecord) error
}

type PhoneRepository interface {
	Find(number string) (*model.Phone, error)
}
