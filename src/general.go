package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// converts string array to golang time var
func convert(raw []string) []time.Time {
	var dates []time.Time
	for _, ts := range raw {
		parsed, _ := time.Parse("2000-01-15", ts)
		dates = append(dates, parsed)
	}
	return dates
}

func clear_screen() {
	out, err := exec.Command("clear").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
