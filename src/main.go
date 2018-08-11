package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Fetch stock data
func getStocks(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func main() {
	args := os.Args[1:] // ignore executable

	url, err := generateURL(args)
	checkErr(err)

	data, err := getStocks(url)
	checkErr(err)

	if len(args) == 1 {
		date, price := parse_TIME_SERIES_DAILY(data)
		fmt.Println(date)
		fmt.Println(price)
	} else {
		mapping := parse_BATCH_STOCK_QUOTES(data)
		fmt.Println(mapping)
	}
}
