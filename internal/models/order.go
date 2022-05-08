package models

import (
	"github.com/shopspring/decimal"
	"time"
)

const (
	StatusOrderNew        = "NEW"
	StatusOrderProcessing = "PROCESSING"
	StatusOrderInvalid    = "INVALID"
	StatusOrderProcessed  = "PROCESSED"
)

type Order struct {
	Number     int
	Status     string
	Accrual    decimal.Decimal
	UploadedAt time.Time
}
