package main

import (
	"os"
)

// Remove all flags from args
// return new args containing only company names
func removeFlags() {

}

func main() {
	args := os.Args[1:] // ignore executable

	// flags, args, err := removeFlags(args)
	// checkErr(err)

	if len(args) == 1 {
		print_TIME_SERIES_DAILY(args)
	} else {
		print_BATCH_STOCK_QUOTES(args)
	}
}
