package http_test

import (
	"encoding/json"
	"errors"
	"goclean/pkg/domain"
	"goclean/pkg/domain/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	ExchangeRateHttp "goclean/pkg/exchange_rate/delivery/http"
)

func TestGetExchangeRateByDateSuccess(t *testing.T) {
	var mockExchangeRate domain.ExchangeRate
	err := faker.FakeData(&mockExchangeRate)
	assert.NoError(t, err)
	mockUseCase := new(mocks.ExchangeRateUsecase)
	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)
	mockUseCase.On("GetExchangeRateByDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockListExchangeRate, nil)

	e := echo.New()
	startDate := "2021-10-17"
	endDate := "2021-10-19"
	req, err := http.NewRequest(echo.GET, "/kurs?startdate="+startDate+"&enddate="+endDate, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("kurs")
	c.SetParamNames("startdate")
	c.SetParamValues(startDate)
	c.SetParamNames("enddate")
	c.SetParamValues(endDate)
	handler := ExchangeRateHttp.ExchangeRateHandler{
		ExchangeRateUsecase: mockUseCase,
	}
	err = handler.GetExchangeRateByDate(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetExchangeRateByDateError(t *testing.T) {
	var mockExchangeRate domain.ExchangeRate
	err := faker.FakeData(&mockExchangeRate)
	assert.NoError(t, err)
	mockUseCase := new(mocks.ExchangeRateUsecase)
	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)
	mockUseCase.On("GetExchangeRateByDate", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, errors.New("invalid params"))

	e := echo.New()
	startDate := "2021-10-17"
	req, err := http.NewRequest(echo.GET, "/kurs?startdate="+startDate, strings.NewReader(""))
	assert.Nil(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("kurs")
	c.SetParamNames("startdate")
	c.SetParamValues(startDate)
	handler := ExchangeRateHttp.ExchangeRateHandler{
		ExchangeRateUsecase: mockUseCase,
	}
	err = handler.GetExchangeRateByDate(c)
	require.Nil(t, err)

	mockUseCase.AssertExpectations(t)
}

func TestGetExchangeRateByCurrencySuccess(t *testing.T) {
	var mockExchangeRate domain.ExchangeRate
	err := faker.FakeData(&mockExchangeRate)
	assert.NoError(t, err)
	mockUseCase := new(mocks.ExchangeRateUsecase)
	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)
	mockUseCase.On("GetExchangeRateByCurrency", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockListExchangeRate, nil)

	e := echo.New()
	startDate := "2021-10-17"
	endDate := "2021-10-19"
	symbol := "USD"
	req, err := http.NewRequest(echo.GET, "/kurs/"+symbol+"?startdate="+startDate+"&enddate="+endDate, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("kurs/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues(symbol)
	c.SetParamNames("startdate")
	c.SetParamValues(startDate)
	c.SetParamNames("enddate")
	c.SetParamValues(endDate)
	handler := ExchangeRateHttp.ExchangeRateHandler{
		ExchangeRateUsecase: mockUseCase,
	}
	err = handler.GetExchangeRateByCurrency(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetExchangeRateByCurrencyError(t *testing.T) {
	var mockExchangeRate domain.ExchangeRate
	err := faker.FakeData(&mockExchangeRate)
	assert.NoError(t, err)
	mockUseCase := new(mocks.ExchangeRateUsecase)
	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)
	mockUseCase.On("GetExchangeRateByCurrency", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, errors.New("invalid params"))

	e := echo.New()
	startDate := "2021-10-17"
	symbol := "USD"
	req, err := http.NewRequest(echo.GET, "/kurs/"+symbol+"?startdate="+startDate, strings.NewReader(""))
	assert.Nil(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("kurs/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues(symbol)
	c.SetParamNames("startdate")
	c.SetParamValues(startDate)
	handler := ExchangeRateHttp.ExchangeRateHandler{
		ExchangeRateUsecase: mockUseCase,
	}
	err = handler.GetExchangeRateByCurrency(c)
	require.Nil(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestStoreSuccess(t *testing.T) {
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

	tempMockExchangeRate := mockExchangeRate
	mockUseCase := new(mocks.ExchangeRateUsecase)

	j, err := json.Marshal(tempMockExchangeRate)
	assert.NoError(t, err)

	mockUseCase.On("Store", mock.Anything, mock.AnythingOfType("*domain.ExchangeRate")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/kurs", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/kurs")

	handler := ExchangeRateHttp.ExchangeRateHandler{
		ExchangeRateUsecase: mockUseCase,
	}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestUpdateSuccess(t *testing.T) {
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

	tempMockExchangeRate := mockExchangeRate
	mockUseCase := new(mocks.ExchangeRateUsecase)

	j, err := json.Marshal(tempMockExchangeRate)
	assert.NoError(t, err)

	mockUseCase.On("Update", mock.Anything, mock.AnythingOfType("*domain.ExchangeRate")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/kurs", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/kurs")

	handler := ExchangeRateHttp.ExchangeRateHandler{
		ExchangeRateUsecase: mockUseCase,
	}
	err = handler.Update(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestDeleteSuccess(t *testing.T) {
	var mockExchangeRate domain.ExchangeRate
	err := faker.FakeData(&mockExchangeRate)
	assert.NoError(t, err)
	mockUseCase := new(mocks.ExchangeRateUsecase)
	mockListExchangeRate := make([]domain.ExchangeRate, 0)
	mockListExchangeRate = append(mockListExchangeRate, mockExchangeRate)
	mockUseCase.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil)

	e := echo.New()
	date := "2021-10-17"
	req, err := http.NewRequest(echo.DELETE, "/kurs/curr/"+date, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/kurs/curr/:date")
	c.SetParamNames("date")
	c.SetParamValues(date)
	handler := ExchangeRateHttp.ExchangeRateHandler{
		ExchangeRateUsecase: mockUseCase,
	}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}