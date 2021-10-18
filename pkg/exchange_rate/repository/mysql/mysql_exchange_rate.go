package mysql

import (
	"context"
	"database/sql"
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

func (m mysqlExchangeRateRepository) Indexing(ctx context.Context, payload []domain.ExchangeRate) error {
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
				&t.CreatedAt,
				&t.UpdatedAt,
				&t.Currency,
				&t.TtCounterType.Buy,
				&t.TtCounterType.Sell,
				&t.BankNoteType.Buy,
				&t.BankNoteType.Sell,
				&t.ERateType.Buy,
				&t.ERateType.Sell,
			)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}
	return result, nil
}

func (m *mysqlExchangeRateRepository) GetExchangeRateByDate(ctx context.Context, startDate string, endDate string) (resp []domain.ExchangeRate, err error) {
	resp, err = m.fetch(ctx, query.GetExchangeRateByDate, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return
}

func (m mysqlExchangeRateRepository) GetExchangeRateByCurrency(ctx context.Context, currency string, startDate string, endDate string) (resp []domain.ExchangeRate, err error) {
	resp, err = m.fetch(ctx, query.GetExchangeRateByCurrency, currency, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return
}

func (m mysqlExchangeRateRepository) Store(ctx context.Context, payload *domain.ExchangeRate) error {
	panic("implement me")
}

func (m mysqlExchangeRateRepository) Update(ctx context.Context, payload *domain.ExchangeRate) error {
	panic("implement me")
}

func (m mysqlExchangeRateRepository) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}