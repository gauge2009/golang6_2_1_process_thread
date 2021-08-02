package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"


)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")

	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "君不见.")
	//设置单元格样式
	style, err := f.NewStyle(`{
    "font":
    {
        "bold": true,
        "family": "font-family",
        "size": 20,
        "color": "#256271"
    }
}`)
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellStyle("Sheet1", "B1", "B1", style)
	f.SetCellValue("Sheet1", "B1", "hello")


	f.SetCellStyle("Sheet2", "A2", "A2", style)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
