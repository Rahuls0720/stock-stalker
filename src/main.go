package main

import (
	"os"
	"strings"
	// // graphing utility
	// "github.com/Arafatk/glot"
)

// prints stock stalker 'man page'
func help() {

}

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

	for _, flag := range flags {
		switch strings.ToLower(flag) {
		case "print history", "f":
			printHistory()

		case "help", "h":
			help()

		case "interval", "i":

		case "print daily change", "r":

		case "track", "t":

		case "update history", "u":
			updateHistory(companies[0], miscArgs)

		default:
			help()
		}
	}

	// if len(companies) == 1 {
	// 	print_TIME_SERIES_DAILY(companies)
	// } else {
	// 	print_BATCH_STOCK_QUOTES(companies)
	// }
}
