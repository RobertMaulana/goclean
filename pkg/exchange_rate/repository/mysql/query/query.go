package query

const (
	GetExchangeRateByDate = `SELECT id,
       created_at,
       updated_at,
       currency,
       tt_counter_buy,
       tt_counter_sell,
       bank_note_buy,
       bank_note_sell,
       e_rate_buy,
       e_rate_sell
from exchange_rates
where created_at >= ? and created_at <= ?
order by created_at`
)
