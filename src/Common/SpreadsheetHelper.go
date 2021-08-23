package Common

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func SmartGetRowColumnCount(path string, sheet_name string) (len_rows int, len_cell int) {

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	//cell, err := f.GetCellValue("ats_result", "A2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheet_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}

	/// 获得行列数
	len_rows = len(rows)
	//fmt.Println(len_rows)
	for _, row := range rows {

		temp := len(row)
		if temp > len_cell {
			len_cell = temp
		}
	}
	//fmt.Println(len_cell)

	return

}
