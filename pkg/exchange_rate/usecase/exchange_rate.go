package usecase

import (
	"context"
	"goclean/pkg/domain"
	"time"
)

type ExchangeRateUsecase struct {
	ExchangeRateRepo domain.ExchangeRateRepository
	contextTimeout time.Duration
}

func NewExchangeRateUsecase(er domain.ExchangeRateRepository, timeout time.Duration) domain.ExchangeRateUsecase {
	return &ExchangeRateUsecase{
		ExchangeRateRepo: er,
		contextTimeout:   timeout,
	}
}

func (e *ExchangeRateUsecase) Indexing(ctx context.Context, payload []domain.ExchangeRate) error {
	return nil
}

func (e *ExchangeRateUsecase) GetExchangeRateByDate(ctx context.Context, startDate string, endDate string) (resp []domain.ExchangeRate, err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	startDate = startDate + " 00:00:01"
	endDate = endDate + " 23:59:59"

	resp, err = e.ExchangeRateRepo.GetExchangeRateByDate(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return
}

func (e *ExchangeRateUsecase) GetExchangeRateBySingleDate(ctx context.Context, symbol string, date string) (resp []domain.ExchangeRate, err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	resp, err = e.ExchangeRateRepo.GetExchangeRateBySingleDate(ctx, symbol, date)
	if err != nil {
		return nil, err
	}
	return
}

func (e *ExchangeRateUsecase) GetExchangeRateBySingleDateOnly(ctx context.Context, date string) (resp []domain.ExchangeRate, err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	resp, err = e.ExchangeRateRepo.GetExchangeRateBySingleDateOnly(ctx, date)
	if err != nil {
		return nil, err
	}
	return
}

func (e *ExchangeRateUsecase) GetExchangeRateByCurrency(ctx context.Context, currency string, startDate string, endDate string) (resp []domain.ExchangeRate, err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	startDate = startDate + " 00:00:01"
	endDate = endDate + " 23:59:59"

	resp, err = e.ExchangeRateRepo.GetExchangeRateByCurrency(ctx, currency, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return
}

func (e *ExchangeRateUsecase) Store(ctx context.Context, payload *domain.ExchangeRate) (err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	list, err := e.GetExchangeRateBySingleDate(ctx, payload.Symbol, payload.Date)
	if err != nil {
		return
	}

	// if data already exist, skip it
	if len(list) > 0 {
		return nil
	}

	err = e.ExchangeRateRepo.Store(ctx, payload)
	return
}

func (e *ExchangeRateUsecase) Update(ctx context.Context, payload *domain.ExchangeRate) (err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	list, err := e.GetExchangeRateBySingleDate(ctx, payload.Symbol, payload.Date)
	if err != nil {
		return
	}

	// if data already exist, update it
	if len(list) > 0 {
		err = e.ExchangeRateRepo.Update(ctx, payload)
		return
	}

	err = domain.ErrNotFound
	return
}

func (e *ExchangeRateUsecase) Delete(ctx context.Context, date string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	existedExchangeRate, err := e.GetExchangeRateBySingleDateOnly(ctx, date)
	if err != nil {
		return
	}

	if len(existedExchangeRate) < 1 {
		return domain.ErrNotFound
	}

	err = e.ExchangeRateRepo.Delete(ctx, date)
	return
}