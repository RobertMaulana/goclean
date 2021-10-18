package query

const (
	GetExchangeRateByDate = `SELECT id,
       symbol,
       tt_counter_buy,
       tt_counter_sell,
       bank_note_buy,
       bank_note_sell,
       e_rate_buy,
       e_rate_sell,
		date
from exchange_rates
where created_at >= ? and created_at <= ?
order by created_at`

	GetExchangeRateBySingleDate = `SELECT id,
       symbol,
       tt_counter_buy,
       tt_counter_sell,
       bank_note_buy,
       bank_note_sell,
       e_rate_buy,
       e_rate_sell,
		date
from exchange_rates
where symbol = ? and date = ?`

	GetExchangeRateByDateOnly = `SELECT id,
       symbol,
       tt_counter_buy,
       tt_counter_sell,
       bank_note_buy,
       bank_note_sell,
       e_rate_buy,
       e_rate_sell,
		date
from exchange_rates
where date = ?`

	GetExchangeRateByCurrency = `SELECT id,
       symbol,
       tt_counter_buy,
       tt_counter_sell,
       bank_note_buy,
       bank_note_sell,
       e_rate_buy,
       e_rate_sell,
		date
from exchange_rates
where symbol = ? and (created_at >= ? and created_at <= ?)
order by created_at`

	Store = `INSERT exchange_rates
SET symbol=?,
	e_rate_buy=?,
    e_rate_sell=?,
    tt_counter_buy=?,
    tt_counter_sell=?,
    bank_note_buy=?,
    bank_note_sell=?,
    date=?`

	Update = `UPDATE exchange_rates 
SET symbol=?,
	e_rate_buy=?,
    e_rate_sell=?,
    tt_counter_buy=?,
    tt_counter_sell=?,
    bank_note_buy=?,
    bank_note_sell=?,
    date=?
WHERE symbol=? and date=?`

	Delete = `DELETE FROM exchange_rates WHERE date = ?`
)
