package model

import "time"

type UsageRecord struct {
	Id        string
	PhoneId   string
	Timestamp time.Time
	Amount    float64
}

type Phone struct {
	Id      string
	Number  string
	Records []*UsageRecord
}
