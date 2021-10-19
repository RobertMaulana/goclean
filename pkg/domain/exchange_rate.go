package domain

import (
	"context"
)

// ExchangeRate entity
type ExchangeRate struct {
	Symbol string `json:"symbol" validate:"required"`
	ERate ExchangeRateDetail `json:"e_rate" validate:"required"`
	TtCounter ExchangeRateDetail `json:"tt_counter" validate:"required"`
	BankNotes ExchangeRateDetail `json:"bank_notes" validate:"required"`
	Date string `json:"date"`
}

// ExchangeRate detail entity
type ExchangeRateDetail struct {
	Sell float64 `json:"jual" validate:"required"`
	Buy float64 `json:"beli" validate:"required"`
}

type ExchangeRateUsecase interface {
	Indexing(ctx context.Context, url string) (err error)
	GetExchangeRateByDate(ctx context.Context, startDate string, endDate string) (resp []ExchangeRate, err error)
	GetExchangeRateByCurrency(ctx context.Context, symbol string, startDate string, endDate string) (resp []ExchangeRate, err error)
	GetExchangeRateBySingleDate(ctx context.Context, symbol string, date string) (resp []ExchangeRate, err error)
	GetExchangeRateBySingleDateOnly(ctx context.Context, date string) (resp []ExchangeRate, err error)
	Store(ctx context.Context, payload *ExchangeRate) error
	Update(ctx context.Context, payload *ExchangeRate) error
	Delete(ctx context.Context, date string) error
}

// ExchangeRateRepository represent the exchange rate's repository contract
type ExchangeRateRepository interface {
	GetExchangeRateByDate(ctx context.Context, startDate string, endDate string) (resp []ExchangeRate, err error)
	GetExchangeRateByCurrency(ctx context.Context, symbol string, startDate string, endDate string) (resp []ExchangeRate, err error)
	GetExchangeRateBySingleDate(ctx context.Context, symbol string, date string) (resp []ExchangeRate, err error)
	GetExchangeRateBySingleDateOnly(ctx context.Context, date string) (resp []ExchangeRate, err error)
	Store(ctx context.Context, payload *ExchangeRate) error
	Update(ctx context.Context, payload *ExchangeRate) error
	Delete(ctx context.Context, date string) error
}