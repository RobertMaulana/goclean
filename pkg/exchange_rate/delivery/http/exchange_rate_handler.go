package http

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"goclean/pkg/domain"
	validator "gopkg.in/go-playground/validator.v9"
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
	e.DELETE(base + "/kurs/curr/:date", handler.Delete)
	e.GET(base + "/kurs", handler.GetExchangeRateByDate)
	e.POST(base + "/kurs", handler.Store)
	e.PUT(base + "/kurs", handler.Update)
}

func (h *ExchangeRateHandler) Indexing(context echo.Context) error {
	ctx := context.Request().Context()

	err := h.ExchangeRateUsecase.Indexing(ctx)
	if err != nil {
		return context.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Done!"})
}

func (h *ExchangeRateHandler) GetExchangeRateByDate(context echo.Context) error {
	startDate := context.QueryParam("startdate")
	endDate := context.QueryParam("enddate")
	ctx := context.Request().Context()

	listArr, err := h.ExchangeRateUsecase.GetExchangeRateByDate(ctx, startDate, endDate)
	if err != nil {
		return context.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
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
		return context.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, listArr)
}

func isRequestValid(m *domain.ExchangeRate) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *ExchangeRateHandler) Store(context echo.Context) (err error) {
	var exchangeRate domain.ExchangeRate
	err = context.Bind(&exchangeRate)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&exchangeRate); !ok {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := context.Request().Context()
	err = h.ExchangeRateUsecase.Store(ctx, &exchangeRate)
	if err != nil {
		return context.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
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
	if ok, err = isRequestValid(&exchangeRate); !ok {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := context.Request().Context()
	err = h.ExchangeRateUsecase.Update(ctx, &exchangeRate)
	if err != nil {
		return context.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, exchangeRate)
}

func (h *ExchangeRateHandler) Delete(context echo.Context) (err error) {
	date := context.Param("date")
	ctx := context.Request().Context()

	err = h.ExchangeRateUsecase.Delete(ctx, date)
	if err != nil {
		return context.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Item removed"})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
