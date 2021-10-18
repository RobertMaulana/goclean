package usecase_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"goclean/pkg/domain"
	"goclean/pkg/domain/mocks"
	"goclean/pkg/exchange_rate/usecase"
	"testing"
	"time"
)


// TestGetExchangeRateByDate represent usecase test for GetExchangeRateByDate
func TestGetExchangeRateByDate(t *testing.T) {
	mockExchangeRateRepo := new(mocks.ExchangeRateRepository)
	mockExchangeRate := domain.ExchangeRate{
		Currency:      "USD",
		ERateType:     domain.ExchangeRateDetail{
			Buy: 14081.00,
			Sell: 14066.00,
		},
		TtCounterType: domain.ExchangeRateDetail{
			Buy: 14229.00,
			Sell: 13929.00,
		},
		BankNoteType:  domain.ExchangeRateDetail{
			Buy: 13929.00,
			Sell: 14229.00,
		},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)

	t.Run("success", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateByDate", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(mockListExchangeRate, nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)
		startDate := "2021-12-02"
		endDate := "2021-12-02"

		list, err := ucase.GetExchangeRateByDate(context.TODO(), startDate, endDate)

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListExchangeRate))

		mockExchangeRateRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateByDate", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(nil,
				errors.New("unexpected error")).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)
		startDate := "2021-12-02"
		endDate := "2021-12-02"

		list, err := ucase.GetExchangeRateByDate(context.TODO(), startDate, endDate)

		assert.Error(t, err)
		assert.Len(t, list, 0)

		mockExchangeRateRepo.AssertExpectations(t)
	})
}
