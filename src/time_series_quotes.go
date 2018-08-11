package main

import (
	"strconv"
	"strings"
	"time"
)

func parse_TIME_SERIES_DAILY(data string) ([]string, []float64) {
	date := make([]string, 0)
	price := make([]float64, 0)

	for _, line := range strings.Split(data, "\n")[1:] {
		if len(line) > 0 {
			arr := strings.Split(line, ",")
			value, _ := strconv.ParseFloat(arr[2], 64)

			date = append(date, arr[0])
			price = append(price, value)
		}
	}
	return date, price
}

// converts string array to golang time
func convert(raw []string) []time.Time {
	var dates []time.Time
	for _, ts := range raw {
		parsed, _ := time.Parse("2000-01-15", ts)
		dates = append(dates, parsed)
	}
	return dates
}
