package http

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"goclean/pkg/domain"
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
	e.GET(base + "/kurs", handler.GetExchangeRateByDate)
	e.GET(base + "/kurs/", handler.GetExchangeRateByCurrency)
	e.POST(base + "/kurs", handler.Store)
	e.PUT(base + "/kurs", handler.Update)
	e.DELETE(base + "/kurs/", handler.Delete)
}

func (h *ExchangeRateHandler) GetExchangeRateByDate(context echo.Context) error {
	startDate := context.QueryParam("startdate")
	endDate := context.QueryParam("enddate")
	cursor := "12"
	ctx := context.Request().Context()

	listArr, nextCursor, err := h.ExchangeRateUsecase.GetExchangeRateByDate(ctx,cursor, startDate, endDate)
	if err != nil {
		return context.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	context.Response().Header().Set(`X-Cursor`, nextCursor)
	return context.JSON(http.StatusOK, listArr)
}

func (h *ExchangeRateHandler) GetExchangeRateByCurrency(context echo.Context) error {
	panic("implement me")
}

func (h *ExchangeRateHandler) Store(context echo.Context) error {
	panic("implement me")
}

func (h *ExchangeRateHandler) Update(context echo.Context) error {
	panic("implement me")
}

func (h *ExchangeRateHandler) Delete(context echo.Context) error {
	panic("implement me")
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
