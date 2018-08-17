package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parse_BATCH_STOCK_QUOTES(data string) map[string]float64 {
	mapping := make(map[string]float64) // maps stock symbol to CURRENT price

	for _, line := range strings.Split(data, "\n")[1:] {
		if len(line) > 0 {
			arr := strings.Split(line, ",")
			fmt.Println(arr)
			value, _ := strconv.ParseFloat(arr[1], 64)

			mapping[arr[0]] = value
		}
	}
	return mapping
}

func print_BATCH_STOCK_QUOTES(companies []string) {
	symbol, err := setsymbols(companies)
	checkErr(err)

	requestURL := url + "function=BATCH_STOCK_QUOTES" + "&" + "symbols=" + symbol + "&" + datatype + "&" + apikey

	for {
		clear_screen()

		data, err := fetchStockData(requestURL)
		checkErr(err)

		mapping := parse_BATCH_STOCK_QUOTES(data)
		for company, value := range mapping {
			fmt.Printf("%v: %v\n", company, value)
		}

		time.Sleep(10 * time.Second)
	}
}
