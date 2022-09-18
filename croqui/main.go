package main

import (
	"croqui/croqui"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const INPUT_FILE = "xxx"
const OUTPUT_FILE = "Book1.xlsx"

func main() {

	fmt.Println("Reading file:", INPUT_FILE)
	// var dict = readExcel()

	var croquis = 10
	file := excelize.NewFile()

	for i := 0; i < croquis; i++ {
		sheetName := "Croqui " + strconv.FormatInt(int64(i), 10)
		file.NewSheet(sheetName)
		// Example
		croqui.GenerateSeedlings(file, sheetName, 5, 45, 3)
	}
	if err := file.SaveAs(OUTPUT_FILE); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Saved in:", OUTPUT_FILE)
}
