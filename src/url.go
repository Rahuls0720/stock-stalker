package main

import (
	"fmt"
)

const url = "https://www.alphavantage.co/query?"
const datatype = "datatype=csv"
const apikey = "apikey=YW42R3PU0KRA106H"

// supported request functions
const TIME_SERIES_DAILY = "function=TIME_SERIES_DAILY"
const TIME_SERIES_INTRADAY = "function=TIME_SERIES_INTRADAY"
const BATCH_STOCK_QUOTES = "function=BATCH_STOCK_QUOTES"

// Updates 'symbol' request parameter
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

func generateURL(companies []string) (string, error) {
	symbol, err := setsymbols(companies)
	if err != nil {
		return "", err
	}

	if len(companies) == 1 { // single stock
		return url + TIME_SERIES_DAILY + "&" + "symbol=" + symbol + "&" + datatype + "&" + apikey, nil
	} else { // multiple stocks
		return url + BATCH_STOCK_QUOTES + "&" + "symbols=" + symbol + "&" + datatype + "&" + apikey, nil
	}
	return "", nil
}
