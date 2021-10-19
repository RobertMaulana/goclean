package http

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"goclean/pkg/domain"
	"goclean/pkg/helpers"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ExchangeRateHandler struct {
	ExchangeRateUsecase domain.ExchangeRateUsecase
}

func NewExchangeRateHandler(e *echo.Echo, exc domain.ExchangeRateUsecase, base string) {
	handler := &ExchangeRateHandler{ExchangeRateUsecase:exc}
	e.GET(base + "/indexing", handler.Indexing)
	e.GET(base + "/kurs/:symbol", handler.GetExchangeRateByCurrency)
	e.GET(base + "/kurs", handler.GetExchangeRateByDate)
	e.POST(base + "/kurs", handler.Store)
	e.PUT(base + "/kurs", handler.Update)
	e.DELETE(base + "/kurs/curr/:date", handler.Delete)
}

func (h *ExchangeRateHandler) Indexing(context echo.Context) error {
	ctx := context.Request().Context()

	scrappingUrl := viper.GetString(`scrapping_url`)
	err := h.ExchangeRateUsecase.Indexing(ctx, scrappingUrl)
	if err != nil {
		return context.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Done!"})
}

func (h *ExchangeRateHandler) GetExchangeRateByDate(context echo.Context) error {
	startDate := context.QueryParam("startdate")
	endDate := context.QueryParam("enddate")
	ctx := context.Request().Context()

	listArr, err := h.ExchangeRateUsecase.GetExchangeRateByDate(ctx, startDate, endDate)
	if err != nil {
		return context.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, listArr)
}

func (h *ExchangeRateHandler) GetExchangeRateByCurrency(context echo.Context) error {
	currency := context.Param("symbol")
	startDate := context.QueryParam("startdate")
	endDate := context.QueryParam("enddate")
	ctx := context.Request().Context()

	listArr, err := h.ExchangeRateUsecase.GetExchangeRateByCurrency(ctx, currency, startDate, endDate)
	if err != nil {
		return context.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, listArr)
}

func (h *ExchangeRateHandler) Store(context echo.Context) (err error) {
	var exchangeRate domain.ExchangeRate
	err = context.Bind(&exchangeRate)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = helpers.IsRequestValid(&exchangeRate); !ok {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := context.Request().Context()
	err = h.ExchangeRateUsecase.Store(ctx, &exchangeRate)
	if err != nil {
		return context.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, exchangeRate)
}

func (h *ExchangeRateHandler) Update(context echo.Context) (err error) {
	var exchangeRate domain.ExchangeRate
	err = context.Bind(&exchangeRate)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = helpers.IsRequestValid(&exchangeRate); !ok {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := context.Request().Context()
	err = h.ExchangeRateUsecase.Update(ctx, &exchangeRate)
	if err != nil {
		return context.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, exchangeRate)
}

func (h *ExchangeRateHandler) Delete(context echo.Context) (err error) {
	date := context.Param("date")
	ctx := context.Request().Context()

	err = h.ExchangeRateUsecase.Delete(ctx, date)
	if err != nil {
		return context.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Item removed"})
}
