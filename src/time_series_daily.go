package main

import (
	"fmt"
	"strings"
	"time"
)

func parse_TIME_SERIES_DAILY(data string) ([]string, []string) {
	date := make([]string, 0)
	price := make([]string, 0)

	for _, line := range strings.Split(data, "\n")[1:] {
		if len(line) > 0 {
			arr := strings.Split(line, ",")
			// value, _ := strconv.ParseFloat(arr[2], 64)
			value := arr[2]

			date = append(date, arr[0])
			price = append(price, value)
		}
	}
	return date, price
}

// func writeToFile(dates []string, prices []string) {
// 	file, err := os.Create("data.txt")
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()

// 	for i, date := range dates {
// 		fmt.Printf("%v: %v\n", date, prices[i])
// 		//file.WriteString(date + ";" + prices[i] + "\n")
// 	}
// }

// gnuplot -p -e "plot 'data.txt' with lines"
func print_TIME_SERIES_DAILY(companies []string) {
	symbol, err := setsymbols(companies)
	checkErr(err)

	requestURL := url + "function=TIME_SERIES_DAILY" + "&" + "symbol=" + symbol + "&" + datatype + "&" + apikey

	for {
		clear_screen()

		data, err := fetchStockData(requestURL)
		checkErr(err)

		dates, prices := parse_TIME_SERIES_DAILY(data)
		for i := 0; i < len(dates); i++ {
			fmt.Printf("%v, %v\n", dates[i], prices[i])
		}
		// 	writeToFile(dates, prices)

		time.Sleep(10 * time.Minute)
	}
}
