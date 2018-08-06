package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	stocks, err := getStocks(url)
	checkErr(err)

	fmt.Println(stocks)
}
