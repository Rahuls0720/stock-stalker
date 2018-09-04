package main

import (
	"fmt"
	"gopkg.in/Iwark/spreadsheet.v2"
	"os"
	"text/tabwriter" // for clean printing
	"errors"
)

// update column of next avaliable row
func updateCell(sheet *spreadsheet.Sheet, column int, data string) {
	sheet.Update(len(sheet.Rows), column, data)

	// make sure changes are reflected
	err := sheet.Synchronize()
	checkErr(err)
}

// // update the sell price and calculates the gain of an existing row
// func updateRow(sheet *spreadsheet.Sheet, company string, purchase string, sell string) { // todo: partial sell of shares
// 	if company == "" || purchase == "" || sell == "" {
// 		return
// 	}

// 	for i := 0; i < len(sheet.Rows); i++ {
// 		row := sheet.Rows[i]

// 		company_name = row[0].Value // todo: hard coded
// 		purchase_price = row[1].Value

// 		if company_name == company && purchase_price == purchase {
// 			updateCell(sheet, 3, sell)
// 			return
// 		}
// 	}
// 	//addRow(company, purchase, sell)
// }

// returns the next avaliable row
// len(sheets.Rows) doesn't necessarily give the next blank row
func nextAvaliableRow(sheet *spreadsheet.Sheet) int, error {
	for i := 0; i < len(sheet.Rows); i++ {
		row := sheet.Rows[i]

		for j := 0; j < len(row); j++ {
			if row[j].Value == "" {
				return i, nil
			}
		}
	}
	return -1, errors.New("nextAvaliableRow() error: no avaliable rows found")
}

// todo: support multiple entires
func updateHistory(company string, miscArgs []string) {
	service, err := spreadsheet.NewService()
	// https://docs.google.com/spreadsheets/d/10YBZtIdH1a1rISx5Uy_J3gqPv7bhVuoZZ78xz9y6f7g/edit#gid=0
	spreadsheet, err := service.FetchSpreadsheet("10YBZtIdH1a1rISx5Uy_J3gqPv7bhVuoZZ78xz9y6f7g")
	checkErr(err)

	sheet, err := spreadsheet.SheetByIndex(0)
	checkErr(err)

	row := nextAvaliableRow(sheet)
	sheet.Update(1, 0, "ar")

	// col := 0
	// for _, arg := range miscArgs {
	// 	sheet.Update(1, ++col, arg)
	// }

	// make sure changes are reflected
	err = sheet.Synchronize()
	checkErr(err)
}

func printHistory() {
	service, err := spreadsheet.NewService()
	// https://docs.google.com/spreadsheets/d/1JxPrXs4bdIV9gGme6gZTBEHRyHGF_OOatOJ1W5cCbFg/edit#gid=0
	spreadsheet, err := service.FetchSpreadsheet("1JxPrXs4bdIV9gGme6gZTBEHRyHGF_OOatOJ1W5cCbFg")
	checkErr(err)

	sheet, err := spreadsheet.SheetByIndex(0)
	checkErr(err)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	for i := 0; i < len(sheet.Rows); i++ {
		row := sheet.Rows[i]

		for j := 0; j < len(row); j++ {
			if j == 1 || j == 3 { // weird formatting issue
				fmt.Fprintf(w, "    ")
			}

			cell := row[j]
			fmt.Fprintf(w, "%v\t", cell.Value)
		}
		fmt.Fprintf(w, "\n")
	}
}
