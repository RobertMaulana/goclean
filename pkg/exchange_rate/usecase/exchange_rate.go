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

func (e *ExchangeRateUsecase) GetExchangeRateByDate(ctx context.Context, cursor string, startDate string, endDate string) (resp []domain.ExchangeRate, nextCursor string, err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	resp, nextCursor, err = e.ExchangeRateRepo.GetExchangeRateByDate(ctx, cursor, startDate, endDate)
	if err != nil {
		return nil, "", err
	}
	return
}

func (e *ExchangeRateUsecase) GetExchangeRateByCurrency(ctx context.Context, cursor string, currency string, startDate string, endDate string) (resp []domain.ExchangeRate, nextCursor string, err error) {
	return nil, "", nil
}

func (e *ExchangeRateUsecase) Store(ctx context.Context, payload *domain.ExchangeRate) error {
	return nil
}

func (e *ExchangeRateUsecase) Update(ctx context.Context, payload *domain.ExchangeRate) error {
	return nil
}

func (e *ExchangeRateUsecase) Delete(ctx context.Context, id int64) error {
	return nil
}