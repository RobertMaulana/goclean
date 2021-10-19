package helpers

import (
	"github.com/sirupsen/logrus"
	"goclean/pkg/domain"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
	"time"
)

type Helpers struct {}

func DateTimeConvert(payload string) (date string, err error) {
	dateSplit := strings.Split(payload, " ")
	yearSplit := strings.Split(dateSplit[2], "")
	payload = dateSplit[0] + " " + monthConvert(dateSplit[1]) + " " + (yearSplit[2]+yearSplit[3]) + " " + strings.ReplaceAll(dateSplit[3], ".", ":") + " WIB"
	dt, err := time.Parse(time.RFC822, payload)
	if err !=  nil {
		return "", err
	}
	return strings.Split(dt.String(), " ")[0], nil
}

func monthConvert(payload string) (month string) {
	switch payload {
	case "Januari":
		return "Jan"
	case "Februari":
		return "Feb"
	case "Maret":
		return "March"
	case "April":
		return "Apr"
	case "Mei":
		return "May"
	case "Juni":
		return "Jun"
	case "Juli":
		return "Jul"
	case "Agustus":
		return "Aug"
	case "September":
		return "Sep"
	case "Oktober":
		return "Oct"
	case "November":
		return "Nov"
	case "Desember":
		return "Dec"
	default:
		return payload
	}
}

func GetStatusCode(err error) int {
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

func IsRequestValid(m *domain.ExchangeRate) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}