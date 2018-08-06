package main

import (
	"fmt"
	"strings"
)

// // request parameters: https://www.alphavantage.co/documentation/#batchquotes
// type Request struct {
// 	url      string // const
// 	function string
// 	symbol   string
// 	interval string // intraday only
// 	apikey   string // const
// }

const url = "https://www.alphavantage.co/query?"
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

	symbol := "symbol="
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

	if len(companies) == 1 {
		return url + TIME_SERIES_DAILY + "&" + symbol + "&" + apikey, nil
	} else { // multiple stocks
		symbol = strings.Replace(symbol, "symbol=", "symbols=", -1)
		return url + BATCH_STOCK_QUOTES + "&" + symbol + "&" + apikey, nil
	}
	return "", nil
}
