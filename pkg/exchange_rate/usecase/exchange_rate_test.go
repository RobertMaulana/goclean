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


// TestScrapping represent usecase test for Scrapping
func TestScrapping(t *testing.T) {
	mockExchangeRateRepo := new(mocks.ExchangeRateRepository)
	mockExchangeRate := domain.ExchangeRate{
			Symbol:      "SGD",
			ERate:     domain.ExchangeRateDetail{
				Buy: 14081.00,
				Sell: 14066.00,
			},
			TtCounter: domain.ExchangeRateDetail{
				Buy: 14229.00,
				Sell: 13929.00,
			},
			BankNotes:  domain.ExchangeRateDetail{
				Buy: 13929.00,
				Sell: 14229.00,
			},
			Date: "2021-10-18",
		}


	//mockExchangeRateFail := domain.ExchangeRate{
	//	Symbol:      "USD",
	//}

	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)

	t.Run("success-skip-existing-data", func(t *testing.T) {
		//tempMockExchangeRate := mockExchangeRate
		mockExchangeRateRepo.On("GetExchangeRateBySingleDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockListExchangeRate, nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)

		err := ucase.Indexing(context.TODO())

		assert.Nil(t, err)
		//assert.Equal(t, mockExchangeRate.Symbol, tempMockExchangeRate.Symbol)

		mockExchangeRateRepo.AssertExpectations(t)
	})

	//t.Run("success-store-new-data", func(t *testing.T) {
	//	mockExchangeRateRepo.On("GetExchangeRateBySingleDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]domain.ExchangeRate{}, nil).Once()
	//	mockExchangeRateRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.ExchangeRate")).Return(nil).Once()
	//
	//	ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)
	//
	//	err := ucase.Indexing(context.TODO())
	//
	//	assert.NoError(t, err)
	//
	//	mockExchangeRateRepo.AssertExpectations(t)
	//})
	//
	//t.Run("error-bad-params", func(t *testing.T) {
	//	mockExchangeRateRepo.On("GetExchangeRateBySingleDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]domain.ExchangeRate{}, nil).Once()
	//	mockExchangeRateRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.ExchangeRate")).Return(domain.ErrBadParamInput).Once()
	//
	//	ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)
	//
	//	err := ucase.Indexing(context.TODO())
	//
	//	assert.Error(t, err)
	//
	//	mockExchangeRateRepo.AssertExpectations(t)
	//})
}

// TestGetExchangeRateByDate represent usecase test for GetExchangeRateByDate
func TestGetExchangeRateByDate(t *testing.T) {
	mockExchangeRateRepo := new(mocks.ExchangeRateRepository)
	mockExchangeRate := domain.ExchangeRate{
		Symbol:      "USD",
		ERate:     domain.ExchangeRateDetail{
			Buy: 14081.00,
			Sell: 14066.00,
		},
		TtCounter: domain.ExchangeRateDetail{
			Buy: 14229.00,
			Sell: 13929.00,
		},
		BankNotes:  domain.ExchangeRateDetail{
			Buy: 13929.00,
			Sell: 14229.00,
		},
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

	t.Run("success-if-data-empty", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateByDate", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return([]domain.ExchangeRate{}, nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)
		startDate := "2021-12-02"
		endDate := "2021-12-02"

		_, err := ucase.GetExchangeRateByDate(context.TODO(), startDate, endDate)

		assert.NoError(t, err)

		mockExchangeRateRepo.AssertExpectations(t)
	})

	t.Run("error-unexpected", func(t *testing.T) {
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

// TestGetExchangeRateByCurrency represent usecase test for GetExchangeRateByCurrency
func TestGetExchangeRateByCurrency(t *testing.T) {
	mockExchangeRateRepo := new(mocks.ExchangeRateRepository)
	mockExchangeRate := domain.ExchangeRate{
		Symbol:      "USD",
		ERate:     domain.ExchangeRateDetail{
			Buy: 14081.00,
			Sell: 14066.00,
		},
		TtCounter: domain.ExchangeRateDetail{
			Buy: 14229.00,
			Sell: 13929.00,
		},
		BankNotes:  domain.ExchangeRateDetail{
			Buy: 13929.00,
			Sell: 14229.00,
		},
	}

	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)

	t.Run("success", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateByCurrency", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockListExchangeRate, nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)
		startDate := "2021-12-02"
		endDate := "2021-12-02"
		currency := "USD"

		list, err := ucase.GetExchangeRateByCurrency(context.TODO(), currency, startDate, endDate)

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListExchangeRate))

		mockExchangeRateRepo.AssertExpectations(t)
	})

	t.Run("success-if-data-empty", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateByCurrency", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]domain.ExchangeRate{}, nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)
		startDate := "2021-12-02"
		endDate := "2021-12-02"
		currency := "USD"

		_, err := ucase.GetExchangeRateByCurrency(context.TODO(), currency, startDate, endDate)

		assert.Nil(t, err)

		mockExchangeRateRepo.AssertExpectations(t)
	})

	t.Run("error-unexpected", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateByCurrency", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil,
			errors.New("unexpected error")).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)
		startDate := "2021-12-02"
		endDate := "2021-12-02"
		currency := "USD"

		list, err := ucase.GetExchangeRateByCurrency(context.TODO(), currency, startDate, endDate)

		assert.Error(t, err)
		assert.Len(t, list, 0)

		mockExchangeRateRepo.AssertExpectations(t)
	})
}

// TestStore represent usecase test for Store
func TestStore(t *testing.T) {
	mockExchangeRateRepo := new(mocks.ExchangeRateRepository)
	mockExchangeRate := domain.ExchangeRate{
		Symbol:      "USD",
		ERate:     domain.ExchangeRateDetail{
			Buy: 14081.00,
			Sell: 14066.00,
		},
		TtCounter: domain.ExchangeRateDetail{
			Buy: 14229.00,
			Sell: 13929.00,
		},
		BankNotes:  domain.ExchangeRateDetail{
			Buy: 13929.00,
			Sell: 14229.00,
		},
		Date: "2021-10-17",
	}

	mockExchangeRateFail := domain.ExchangeRate{
		Symbol:      "USD",
	}

	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)

	t.Run("success-skip-existing-data", func(t *testing.T) {
		tempMockExchangeRate := mockExchangeRate
		mockExchangeRateRepo.On("GetExchangeRateBySingleDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockListExchangeRate, nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)

		err := ucase.Store(context.TODO(), &mockExchangeRate)

		assert.Nil(t, err)
		assert.Equal(t, mockExchangeRate.Symbol, tempMockExchangeRate.Symbol)

		mockExchangeRateRepo.AssertExpectations(t)
	})

	t.Run("success-store-new-data", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateBySingleDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]domain.ExchangeRate{}, nil).Once()
		mockExchangeRateRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.ExchangeRate")).Return(nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)

		err := ucase.Store(context.TODO(), &mockExchangeRate)

		assert.NoError(t, err)

		mockExchangeRateRepo.AssertExpectations(t)
	})

	t.Run("error-bad-params", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateBySingleDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]domain.ExchangeRate{}, nil).Once()
		mockExchangeRateRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.ExchangeRate")).Return(domain.ErrBadParamInput).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)

		err := ucase.Store(context.TODO(), &mockExchangeRateFail)

		assert.Error(t, err)

		mockExchangeRateRepo.AssertExpectations(t)
	})
}

// TestUpdate represent usecase test for Update
func TestUpdate(t *testing.T) {
	mockExchangeRateRepo := new(mocks.ExchangeRateRepository)
	mockExchangeRate := domain.ExchangeRate{
		Symbol:      "USD",
		ERate:     domain.ExchangeRateDetail{
			Buy: 14081.00,
			Sell: 14066.00,
		},
		TtCounter: domain.ExchangeRateDetail{
			Buy: 14229.00,
			Sell: 13929.00,
		},
		BankNotes:  domain.ExchangeRateDetail{
			Buy: 13929.00,
			Sell: 14229.00,
		},
		Date: "2021-10-17",
	}

	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)

	t.Run("success-update-if-data-exist", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateBySingleDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockListExchangeRate, nil).Once()
		mockExchangeRateRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.ExchangeRate")).Return(nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)

		err := ucase.Update(context.TODO(), &mockExchangeRate)

		assert.NoError(t, err)

		mockExchangeRateRepo.AssertExpectations(t)
	})

	t.Run("error-if-data-empty", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateBySingleDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]domain.ExchangeRate{}, nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)

		err := ucase.Update(context.TODO(), &mockExchangeRate)

		assert.Error(t, err)

		mockExchangeRateRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockExchangeRateRepo := new(mocks.ExchangeRateRepository)
	mockExchangeRate := domain.ExchangeRate{
		Symbol:      "USD",
		ERate:     domain.ExchangeRateDetail{
			Buy: 14081.00,
			Sell: 14066.00,
		},
		TtCounter: domain.ExchangeRateDetail{
			Buy: 14229.00,
			Sell: 13929.00,
		},
		BankNotes:  domain.ExchangeRateDetail{
			Buy: 13929.00,
			Sell: 14229.00,
		},
		Date: "2021-10-18",
	}

	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)

	t.Run("success", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateBySingleDateOnly", mock.Anything, mock.AnythingOfType("string")).Return(mockListExchangeRate, nil).Once()
		mockExchangeRateRepo.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)

		err := ucase.Delete(context.TODO(), mockExchangeRate.Date)

		assert.NoError(t, err)
		mockExchangeRateRepo.AssertExpectations(t)
	})

	t.Run("error-if-data-not-exist", func(t *testing.T) {
		mockExchangeRateRepo.On("GetExchangeRateBySingleDateOnly", mock.Anything, mock.AnythingOfType("string")).Return([]domain.ExchangeRate{}, nil).Once()

		ucase := usecase.NewExchangeRateUsecase(mockExchangeRateRepo, time.Second * 2)

		err := ucase.Delete(context.TODO(), mockExchangeRate.Date)

		assert.Error(t, err)
		mockExchangeRateRepo.AssertExpectations(t)
	})
}
