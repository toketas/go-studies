package croqui

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xuri/excelize/v2"
)

func GenerateSeedlings(excel *excelize.File, sheetName string, experiments int, columns int, blocks int) {
	fmt.Printf("Running with %d columns/%d experiments\n", columns, experiments)

	const SEEDLINGS = 5
	var l = 0
	var c = 0

	for i := 0; i < blocks; i++ {
		rand.Seed(time.Now().UnixNano())
		var array = rand.Perm(experiments)

		for j := 0; j < experiments; j++ {
			var value = array[j]
			for k := 0; k < SEEDLINGS; k++ {
				// sheet.write(excel, sheetName, l, c, value, i%2==0)
				fmt.Printf("%d ", value)
				c++
				if c == columns-1 {
					l++
					c = 0
				}
			}
		}
	}
}
