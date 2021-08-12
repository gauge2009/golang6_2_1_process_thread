package mainExcelize

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	//"github.com/xuri/excelize/v2"
	"strconv"
)

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func main() {
	/// 计算值
	//SmartCheckCalc("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.xlsx")

	//
	///// 预期值
	//SmartCheckCalc("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.expect.xlsx")

	//path_root := getCurrentAbPathByExecutable() // go run 与 go build 执行不一样，前者使用C:\Users\Administrator\AppData\Local\Temp
	path_root := getCurrentAbPath() // // go run 与 go build  统一使用go run 制指定的
	fmt.Println("getCurrentAbPath = ", path_root)

	sheet_name := "ats_result_viewmodel"
	//
	///// 获得行列数
	//path_calc :="D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.xlsx"
	//path_expect :="D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.expect.xlsx"
	path_calc := path_root + "\\inspect_data.xlsx"
	path_expect := path_root + "\\inspect_data.expect.xlsx"

	len_rows, len_cell := SmartGetRowColumnCount(path_calc, sheet_name)
	fmt.Println(len_rows)
	fmt.Println(len_cell)

	//column_dictionary := make(map [int] string)
	column_dictionary := map[int]string{
		0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J", 10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R", 18: "S", 19: "T", 20: "U", 21: "V", 22: "W", 23: "X", 24: "Y", 25: "Z",
		26: "AA", 27: "AB", 28: "AC", 29: "AD", 30: "AE", 31: "AF", 32: "AG", 33: "AH", 34: "AI", 35: "AJ", 36: "AK", 37: "AL", 38: "AM", 39: "AN", 40: "AO", 41: "AP", 42: "AQ", 43: "AR", 44: "AS", 45: "AT", 46: "AU", 47: "AV", 48: "AW", 49: "AX", 50: "AY", 51: "AZ",
		52: "BA", 53: "BB", 54: "BC", 55: "BD", 56: "BE", 57: "BF", 58: "BG", 59: "BH", 60: "BI", 61: "BJ", 62: "BK", 63: "BL", 64: "BM", 65: "BN", 66: "BO", 67: "BP", 68: "BQ", 69: "BR", 70: "BS", 71: "BT", 72: "BU", 73: "BV", 74: "BW", 75: "BX", 76: "BY", 77: "BZ",
		78: "CA", 79: "CB", 80: "CC", 81: "CD", 82: "CE", 83: "CF", 84: "CG", 85: "CH", 86: "CI", 87: "CJ", 88: "CK", 89: "CL", 90: "CM", 91: "CN", 92: "CO", 93: "CP", 94: "CQ", 95: "CR", 96: "CS", 97: "CT", 98: "CU", 99: "CV", 100: "CW", 101: "CX", 102: "CY", 103: "CZ",
	}

	SmartCompareThose(path_calc, path_expect, sheet_name, len_rows, len_cell, column_dictionary)

	//SmartGetCols(path_calc,sheet_name)
}

func SmartCheckCalc(path string) {
	//arg := []string{}// 切片
	//arg = append(arg,"inspect,60004")// 切片后加一个元素
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("ats_result_viewmodel", "A2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("ats_result_viewmodel")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}

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

func SmartCompareThose(path_calc string, path_expect string, sheet_name string, len_rows int, len_cell int, column_dictionary map[int]string) {

	//arg := []string{}// 切片
	//arg = append(arg,"inspect,60004")// 切片后加一个元素

	/// 计算值
	f, err := excelize.OpenFile(path_calc)
	if err != nil {
		fmt.Println(err)
		return
	}
	//cell, err := f.GetCellValue("ats_result", "A2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cell)
	rows, err := f.GetRows(sheet_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	/// 预期值
	f_exp, err := excelize.OpenFile(path_expect)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows_exp, err := f_exp.GetRows(sheet_name)
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

	f_diff := excelize.NewFile()
	// Create a new sheet.
	index := f_diff.NewSheet(sheet_name)
	style, err := f_diff.NewStyle(`{
   "font": {  "bold": true,  "family": "font-family",  "size": 20, "color": "#FB5A19"  },
	"fill":{"type":"gradient","color":["#FFFF00","#E0EBF5"],"shading":1},
	"border":[{"type":"left","color":"42CFD2","style":6},{"type":"top","color":"42CFD2","style":6},{"type":"bottom","color":"42CFD2","style":6},{"type":"right","color":"42CFD2","style":6}]
}`)
	err = f_diff.SetHeaderFooter(sheet_name, &excelize.FormatHeaderFooter{
		DifferentFirst:   true,
		DifferentOddEven: true,
		OddHeader:        "&R&P",
		OddFooter:        "&C&F",
		EvenHeader:       "&L&P",
		EvenFooter:       "&L&D&R&T",
		FirstHeader:      `&CCenter &"-,Bold"Bold&"-,Regular"HeaderU+000A&D`,
	})

	///差异文件的行号索引（2基,因第一行已被表头占据）
	index_dif_expect := 2
	index_dif_calc := 2
	for i := 0; i < len_rows; i++ {
		for j := 0; j < len_cell; j++ {
			axis_name := column_dictionary[j] + strconv.Itoa(i+1)
			///先形成表头
			if i == 0 {
				//f_diff.SetCellValue("ats_result", "D2", rows[i][j])
				f_diff.SetCellValue(sheet_name, axis_name, rows_exp[i][j])
			} else {
				/// 记录预期值
				//index_dif_expect =index_dif_expect
				axis_name = column_dictionary[j] + strconv.Itoa(index_dif_expect)
				f_diff.SetCellValue(sheet_name, axis_name, rows_exp[i][j])
				/// 记录差异
				if rows[i][j] != rows_exp[i][j] {
					fmt.Print(rows[i][j], "\t")
					fmt.Print(rows_exp[i][j], "\t")

					/// 记录计算差异值
					index_dif_calc = index_dif_calc + 1
					axis_name_calc := column_dictionary[j] + strconv.Itoa(index_dif_calc)
					f_diff.SetCellValue(sheet_name, axis_name_calc, rows[i][j])
					f_diff.SetCellStyle(sheet_name, axis_name_calc, axis_name_calc, style)
					//f_diff.SetColWidth(sheet_name, axis_name_calc, axis_name_calc, 400)
				}
			}
		}
		index_dif_calc++
		index_dif_expect = index_dif_calc
		fmt.Println()
	}

	f_diff.SetActiveSheet(index)
	//path_root := getCurrentAbPathByExecutable() // go run 与 go build 执行不一样，前者使用C:\Users\Administrator\AppData\Local\Temp
	path_root := getCurrentAbPath() // // go run 与 go build  统一使用go run 制指定的
	fmt.Println("getCurrentAbPath = ", path_root)
	//save_to_path := "D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.compare.xlsx"
	save_to_path := path_root + "\\inspect_data.compare.xlsx"
	if err := f_diff.SaveAs(save_to_path); err != nil {
		fmt.Println(err)
	}

	//
	//	f_diff := excelize.NewFile()
	//	// Create a new sheet.
	//	index := f_diff.NewSheet("ats_result")
	//
	//	// Set value of a cell.
	//	f_diff.SetCellValue("ats_result", "A2", "君不见.")
	//	//设置单元格样式
	//	style, err := f_diff.NewStyle(`{
	//   "font": {  "bold": true,  "family": "font-family",  "size": 20, "color": "#FB5A19"  },
	//	"fill":{"type":"gradient","color":["#FFFF00","#E0EBF5"],"shading":1}
	//}`)
	//     /*  删除线  */
	//	//style, err := f_diff.NewStyle(`{"border":[{"type":"left","color":"0000FF","style":3},{"type":"top","color":"00FF00","style":4},{"type":"bottom","color":"FFFF00","style":5},{"type":"right","color":"FF0000","style":6},{"type":"diagonalDown","color":"A020F0","style":7},{"type":"diagonalUp","color":"A020F0","style":8}]}`)
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//}
	//	/*  渐变背景  */
	//	//style, err := f_diff.NewStyle(`{"fill":{"type":"gradient","color":["#FB5A19","#E0EBF5"],"shading":1}}`)
	//
	//	f_diff.SetCellStyle("ats_result", "A2", "A2", style)
	//	// Set active sheet of the workbook.
	//	f_diff.SetActiveSheet(index)
	//	// Save xlsx file by the given path.
	//	if err := f_diff.SaveAs("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.compare.xlsx"); err != nil {
	//		fmt.Println(err)
	//	}

}

func SmartGetCols(path string, sheet_name string) {
	/// 计算值
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	cols, err := f.Cols(sheet_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	for cols.Next() {
		col, err := cols.Rows()
		if err != nil {
			fmt.Println(err)
		}
		for _, rowCell := range col {
			fmt.Print(rowCell, "\t")
		}
		fmt.Println()
	}
}
