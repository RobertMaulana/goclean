package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"goclean/pkg/domain"
	"goclean/pkg/exchange_rate/repository/mysql/query"
)

type mysqlExchangeRateRepository struct {
	Conn *sql.DB
}

func NewMysqlExchangeRateRepository(Conn *sql.DB) domain.ExchangeRateRepository {
	return &mysqlExchangeRateRepository{Conn}
}

func (m *mysqlExchangeRateRepository) Indexing(ctx context.Context, payload []domain.ExchangeRate) error {
	panic("implement me")
}

func (m *mysqlExchangeRateRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.ExchangeRate, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.ExchangeRate, 0)
	for rows.Next() {
		t := domain.ExchangeRate{}
		exchangeId := int64(0)
		err = rows.Scan(
				&exchangeId,
				&t.Symbol,
				&t.TtCounter.Buy,
				&t.TtCounter.Sell,
				&t.BankNotes.Buy,
				&t.BankNotes.Sell,
				&t.ERate.Buy,
				&t.ERate.Sell,
				&t.Date,
			)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}
	return result, nil
}

func (m *mysqlExchangeRateRepository) GetExchangeRateBySingleDate(ctx context.Context, symbol string, date string) (resp []domain.ExchangeRate, err error) {
	resp, err = m.fetch(ctx, query.GetExchangeRateBySingleDate, symbol, date)
	if err != nil {
		return nil, err
	}
	return
}

func (m *mysqlExchangeRateRepository) GetExchangeRateByDate(ctx context.Context, startDate string, endDate string) (resp []domain.ExchangeRate, err error) {
	resp, err = m.fetch(ctx, query.GetExchangeRateByDate, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return
}

func (m *mysqlExchangeRateRepository) GetExchangeRateBySingleDateOnly(ctx context.Context, date string) (resp []domain.ExchangeRate, err error) {
	resp, err = m.fetch(ctx, query.GetExchangeRateByDateOnly, date)
	if err != nil {
		return nil, err
	}
	return
}

func (m *mysqlExchangeRateRepository) GetExchangeRateByCurrency(ctx context.Context, currency string, startDate string, endDate string) (resp []domain.ExchangeRate, err error) {
	fmt.Printf("curr %s start %s end %s", currency, startDate, endDate)

	resp, err = m.fetch(ctx, query.GetExchangeRateByCurrency, currency, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return
}

func (m *mysqlExchangeRateRepository) Store(ctx context.Context, payload *domain.ExchangeRate) (err error){
	stmt, err := m.Conn.PrepareContext(ctx, query.Store)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		payload.Symbol,
		payload.ERate.Buy, payload.ERate.Sell,
		payload.TtCounter.Buy, payload.TtCounter.Sell,
		payload.BankNotes.Buy, payload.BankNotes.Sell,
		payload.Date)
	if err != nil {
		return
	}
	return
}

func (m *mysqlExchangeRateRepository) Update(ctx context.Context, payload *domain.ExchangeRate) (err error) {
	stmt, err := m.Conn.PrepareContext(ctx, query.Update)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		payload.Symbol,
		payload.ERate.Buy, payload.ERate.Sell,
		payload.TtCounter.Buy, payload.TtCounter.Sell,
		payload.BankNotes.Buy, payload.BankNotes.Sell,
		payload.Date,
		payload.Symbol, payload.Date)
	if err != nil {
		return
	}
	return
}

func (m *mysqlExchangeRateRepository) Delete(ctx context.Context, date string) (err error) {
	stmt, err := m.Conn.PrepareContext(ctx, query.Delete)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, date)
	if err != nil {
		return
	}

	_, err = res.RowsAffected()
	if err != nil {
		return
	}

	return
}