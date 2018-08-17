package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// alpha avantage API parameters
// generic location for use by both batch_stock_quotes.go and time_series_daily.go
const url = "https://www.alphavantage.co/query?"
const datatype = "datatype=csv"
const apikey = "apikey=YW42R3PU0KRA106H"

// Updates (company) 'symbol' request parameter
func setsymbols(companies []string) (string, error) {
	if len(companies) == 0 {
		return "", fmt.Errorf("Error: no stock symbols provided, must provide at least 1")
	}

	var symbol string
	for _, company := range companies {
		symbol += company + ","
	}
	return symbol[:len(symbol)-1], nil
}

func fetchStockData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// TODO: if body says [    "Information": "Thank you for using Alpha Vantage! Please visit https://www.alphavantage.co/premium/ if you would like to have a higher API call volume."] retry
	return string(body), nil
}
