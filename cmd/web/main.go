package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"time"

	exchangeRateRepo "goclean/pkg/exchange_rate/repository/mysql"
	exchangeRateUsecase "goclean/pkg/exchange_rate/usecase"
	exchangeRateHandler "goclean/pkg/exchange_rate/delivery/http"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal()
		}
	}()

	ech := echo.New()
	excRateRepo := exchangeRateRepo.NewMysqlExchangeRateRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	excRateUsecase := exchangeRateUsecase.NewExchangeRateUsecase(excRateRepo, timeoutContext)
	exchangeRateHandler.NewExchangeRateHandler(ech, excRateUsecase, viper.GetString("base.api"))

	log.Fatal(ech.Start(viper.GetString("server.address")))
}