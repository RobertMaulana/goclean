// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "goclean/pkg/domain"

	mock "github.com/stretchr/testify/mock"
)

// ExchangeRateUsecase is an autogenerated mock type for the ExchangeRateUsecase type
type ExchangeRateUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ExchangeRateUsecase) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetExchangeRateByCurrency provides a mock function with given fields: ctx, cursor, currency, startDate, endDate
func (_m *ExchangeRateUsecase) GetExchangeRateByCurrency(ctx context.Context, cursor string, currency string, startDate string, endDate string) ([]domain.ExchangeRate, string, error) {
	ret := _m.Called(ctx, cursor, currency, startDate, endDate)

	var r0 []domain.ExchangeRate
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) []domain.ExchangeRate); ok {
		r0 = rf(ctx, cursor, currency, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ExchangeRate)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) string); ok {
		r1 = rf(ctx, cursor, currency, startDate, endDate)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string, string) error); ok {
		r2 = rf(ctx, cursor, currency, startDate, endDate)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetExchangeRateByDate provides a mock function with given fields: ctx, cursor, startDate, endDate
func (_m *ExchangeRateUsecase) GetExchangeRateByDate(ctx context.Context, cursor string, startDate string, endDate string) ([]domain.ExchangeRate, string, error) {
	ret := _m.Called(ctx, cursor, startDate, endDate)

	var r0 []domain.ExchangeRate
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) []domain.ExchangeRate); ok {
		r0 = rf(ctx, cursor, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ExchangeRate)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) string); ok {
		r1 = rf(ctx, cursor, startDate, endDate)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string) error); ok {
		r2 = rf(ctx, cursor, startDate, endDate)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Indexing provides a mock function with given fields: ctx, payload
func (_m *ExchangeRateUsecase) Indexing(ctx context.Context, payload []domain.ExchangeRate) error {
	ret := _m.Called(ctx, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.ExchangeRate) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: ctx, payload
func (_m *ExchangeRateUsecase) Store(ctx context.Context, payload *domain.ExchangeRate) error {
	ret := _m.Called(ctx, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ExchangeRate) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, payload
func (_m *ExchangeRateUsecase) Update(ctx context.Context, payload *domain.ExchangeRate) error {
	ret := _m.Called(ctx, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ExchangeRate) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
