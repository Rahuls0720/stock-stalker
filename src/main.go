package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
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

// BATCH_STOCK_QUOTES
func parse(companies []string, data string) map[string]float64 {
	mapping := make(map[string]float64) // maps stock symbol to price

	regex, _ := regexp.Compile("price(.*?)[0-9]+.[0-9]+")
	for i, json := range regex.FindAllString(data, -1) {

		regex1, _ := regexp.Compile("[0-9]+.[0-9]+")
		price, _ := strconv.ParseFloat(regex1.FindString(json), 64)
		mapping[companies[i]] = price
	}
	return mapping
}

// TIME_SERIES_DAILY
func parse1(data string) ([]string, []float64) {
	date := make([]string, 0)
	price := make([]float64, 0)

	for _, line := range strings.Split(data, "\n") {
		if len(line) > 0 {
			arr := strings.Split(line, ",")
			value, _ := strconv.ParseFloat(arr[2], 64)

			date = append(date, arr[0])
			price = append(price, value)
		}
	}
	return date, price
}

func main() {
	args := os.Args[1:] // ignore executable

	url, err := generateURL(args)
	checkErr(err)

	data, err := getStocks(url)
	checkErr(err)
	//fmt.Println(data)

	date, price := parse1(data)
	fmt.Println(date)
	fmt.Println(price)
}
