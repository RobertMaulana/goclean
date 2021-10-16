package domain

import (
	"context"
	"time"
)

// ExchangeRate entity
type ExchangeRate struct {
	ID int64 `json:"id"`
	Currency string `json:"currency" validate:"required"`
	ERateType ExchangeRateDetail `json:"e_rate_type" validate:"required"`
	TtCounterType ExchangeRateDetail `json:"tt_counter_sell" validate:"required"`
	BankNoteType ExchangeRateDetail `json:"bank_note_sell" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ExchangeRate detail entity
type ExchangeRateDetail struct {
	Sell float64 `json:"sell" validate:"required"`
	Buy float64 `json:"buy" validate:"required"`
}

type ExchangeRateUsecase interface {
	Indexing(ctx context.Context, payload []ExchangeRate) error
	GetExchangeRateByDate(ctx context.Context, cursor string, startDate string, endDate string) (resp []ExchangeRate, nextCursor string, err error)
	GetExchangeRateByCurrency(ctx context.Context, cursor string, currency string, startDate string, endDate string) (resp []ExchangeRate, nextCursor string, err error)
	Store(ctx context.Context, payload *ExchangeRate) error
	Update(ctx context.Context, payload *ExchangeRate) error
	Delete(ctx context.Context, id int64) error
}

// ExchangeRateRepository represent the exchange rate's repository contract
type ExchangeRateRepository interface {
	Indexing(ctx context.Context, payload []ExchangeRate) error
	GetExchangeRateByDate(ctx context.Context, cursor string, startDate string, endDate string) (resp []ExchangeRate, nextCursor string, err error)
	GetExchangeRateByCurrency(ctx context.Context, cursor string, currency string, startDate string, endDate string) (resp []ExchangeRate, nextCursor string, err error)
	Store(ctx context.Context, payload *ExchangeRate) error
	Update(ctx context.Context, payload *ExchangeRate) error
	Delete(ctx context.Context, id int64) error
}