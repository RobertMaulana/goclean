package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"goclean/pkg/domain"
	"goclean/pkg/helpers"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func (e *ExchangeRateUsecase) Indexing(ctx context.Context, url string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]domain.ExchangeRate, 0)

	doc.Find("tbody").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(domain.ExchangeRate)
		row.Symbol = sel.Find(".sticky-col p").Text()

		sel.Find("p").Each(func(i int, s *goquery.Selection) {
			doc.Find(".o-kurs-refresh-description").Each(func(i int, sel *goquery.Selection) {
				date, err := helpers.DateTimeConvert(sel.Find(".refresh-date").Text() + " WIB")
				if err != nil {
					logrus.Error(err)
				}
				row.Date = date
			})
			href, ok := s.Attr("rate-type")
			if ok {
				if href == "ERate-buy" {
					ERateBuy, err := strconv.ParseFloat(
						strings.ReplaceAll(strings.ReplaceAll(s.Text(), ".", ""), ",", ".") ,
						64)
					if err != nil {
						log.Println(err)
					}
					row.ERate.Buy = ERateBuy
				}
				if href == "ERate-sell" {
					ERateSell, err := strconv.ParseFloat(
						strings.ReplaceAll(strings.ReplaceAll(s.Text(), ".", ""), ",", ".") ,
						64)
					if err != nil {
						log.Println(err)
					}
					row.ERate.Sell = ERateSell
				}
				if href == "TT-buy" {
					TTBuy, err := strconv.ParseFloat(
						strings.ReplaceAll(strings.ReplaceAll(s.Text(), ".", ""), ",", ".") ,
						64)
					if err != nil {
						log.Println(err)
					}
					row.TtCounter.Buy = TTBuy
				}
				if href == "TT-sell" {
					TTSell, err := strconv.ParseFloat(
						strings.ReplaceAll(strings.ReplaceAll(s.Text(), ".", ""), ",", ".") ,
						64)
					if err != nil {
						log.Println(err)
					}
					row.TtCounter.Sell = TTSell
				}
				if href == "BN-buy" {
					BNbuy, err := strconv.ParseFloat(
						strings.ReplaceAll(strings.ReplaceAll(s.Text(), ".", ""), ",", ".") ,
						64)
					if err != nil {
						log.Println(err)
					}
					row.BankNotes.Buy = BNbuy
				}
				if href == "BN-sell" {
					BNsell, err := strconv.ParseFloat(
						strings.ReplaceAll(strings.ReplaceAll(s.Text(), ".", ""), ",", ".") ,
						64)
					if err != nil {
						log.Println(err)
					}
					row.BankNotes.Sell = BNsell
				}
			}
		})

		rows = append(rows, *row)
	})
	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	var exchangeRate []domain.ExchangeRate
	err = json.Unmarshal(bts, &exchangeRate)

	for _, value := range exchangeRate {
		err = e.Store(ctx, &value)
	}

	fmt.Println("Done!")
	return
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