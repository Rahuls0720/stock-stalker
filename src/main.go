package main

import (
	"fmt"
	"os"
	"strings"
)

// parse arguments and return flags, companies, and all other args as seperate arrays
func parseArgs(args []string) ([]string, []string, []string) {
	flags := make([]string, 0)
	companies := make([]string, 0)
	miscArgs := make([]string, 0)

	for _, ele := range args {
		if strings.Contains(ele, "-") {
			flagArr := strings.Split(ele, "")
			for _, flag := range flagArr {
				if flag != "-" {
					flags = append(flags, flag)
				}
			}
		} else if !strings.ContainsAny(ele, "0123456789") {
			companies = append(companies, ele)
		} else {
			miscArgs = append(miscArgs, ele)
		}
	}
	return flags, companies, miscArgs
}

func main() {
	flags, companies, miscArgs := parseArgs(os.Args[1:]) // ignore executable

	if len(companies) == 1 {
		print_TIME_SERIES_DAILY(args)
	} else {
		print_BATCH_STOCK_QUOTES(args)
	}
}
