package main

import (
	"strconv"
	"strings"
)

func parse_BATCH_STOCK_QUOTES(data string) map[string]float64 {
	mapping := make(map[string]float64) // maps stock symbol to CURRENT price

	for _, line := range strings.Split(data, "\n")[1:] {
		if len(line) > 0 {
			arr := strings.Split(line, ",")
			value, _ := strconv.ParseFloat(arr[1], 64)

			mapping[arr[0]] = value
		}
	}
	return mapping
}
