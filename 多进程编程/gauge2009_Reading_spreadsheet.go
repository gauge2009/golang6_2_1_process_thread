package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func main() {
	/// 计算值
	//CheckCalc("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.xlsx")

	//
	///// 预期值
	//CheckCalc("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.expect.xlsx")

	//
	///// 获得行列数
	path_calc :="D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.xlsx"
	path_expect :="D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.expect.xlsx"
 	len_rows, len_cell := GetRowColumnCount(path_calc)
	fmt.Println(len_rows)
	fmt.Println(len_cell)

 	//column_dictionary := make(map [int] string)
	column_dictionary :=  map [int] string{
		0:"A",
		1:"B",
		2:"C",
		3:"D",
		4:"E",
		5:"F",
		6:"G",
		7:"H",
		8:"I",
		9:"J",
		10:"K",
		11:"L",
		12:"M",
		13:"N",
		14:"O",
		15:"P",
		16:"Q",
		17:"R",
		18:"S",
		19:"T",
		20:"U",
		21:"V",
		22:"W",
		23:"X",
		24:"Y",
		25:"Z",
		26:"AA",
		27:"AB",
		28:"AC",
		29:"AD",
		30:"AE",
		31:"AF",
		32:"AG",
		33:"AH",
		34:"AI",
		35:"AJ",
		36:"AK",
		37:"AL",
		38:"AM",
		39:"AN",
		40:"AO",
		41:"AP",
		42:"AQ",
		43:"AR",
		44:"AS",
		45:"AT",
		46:"AU",
		47:"AV",
		48:"AW",
		49:"AX",
		50:"AY",
		51:"AZ",
	}


	CompareThose(path_calc,path_expect,len_rows,len_cell,column_dictionary);

	//GetCols(path_calc)
}

func CheckCalc( path string) {
	arg := []string{}// 切片
	arg = append(arg,"inspect,60004")// 切片后加一个元素
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("ats_result", "A2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("ats_result")
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


func GetRowColumnCount( path string)(len_rows int ,len_cell int) {
	arg := []string{}// 切片
	arg = append(arg,"inspect,60004")// 切片后加一个元素
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
	rows, err := f.GetRows("ats_result")
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
		len_cell = len(row)
	}
	//fmt.Println(len_cell)

	return


}


func CompareThose( path_calc string,path_expect string,len_rows int ,len_cell int, column_dictionary map [int] string) {
	arg := []string{}// 切片
	arg = append(arg,"inspect,60004")// 切片后加一个元素

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
	rows, err := f.GetRows("ats_result")
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
	rows_exp, err := f_exp.GetRows("ats_result")
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
	index := f_diff.NewSheet("ats_result")
	style, err := f_diff.NewStyle(`{
   "font": {  "bold": true,  "family": "font-family",  "size": 20, "color": "#FB5A19"  },
	"fill":{"type":"gradient","color":["#FFFF00","#E0EBF5"],"shading":1}
}`)
	///差异文件的行号索引（2基,因第一行已被表头占据）
	index_dif_expect := 2
	index_dif_calc := 2
	for i:=0; i< len_rows;i++{
		for j:=0;j<len_cell;j++{
			axis_name :=  column_dictionary[j]+ strconv.Itoa(i+1)
			///先形成表头
			if i==0 {
				//f_diff.SetCellValue("ats_result", "D2", rows[i][j])
				f_diff.SetCellValue("ats_result", axis_name, rows[i][j])
			}else {
				/// 记录预期值
				index_dif_expect =index_dif_expect
				axis_name = column_dictionary[j] + strconv.Itoa(index_dif_expect)
				f_diff.SetCellValue("ats_result", axis_name, rows[i][j])
				/// 记录差异
				if rows[i][j] != rows_exp[i][j] {
					fmt.Print(rows[i][j], "\t")
					fmt.Print(rows_exp[i][j], "\t")
					///// 记录预期值
					//index_dif =index_dif
					//axis_name = column_dictionary[j] + strconv.Itoa(index_dif)
					//f_diff.SetCellValue("ats_result", axis_name, rows[i][j])
					//f_diff.SetSheetRow("ats_result", "A1", &[]interface{}{rows[i]})
					/// 记录计算差异值
					index_dif_calc =index_dif_calc+1
					axis_name_calc  := column_dictionary[j] + strconv.Itoa(index_dif_calc)
					f_diff.SetCellValue("ats_result", axis_name_calc, rows_exp[i][j])
					f_diff.SetCellStyle("ats_result", axis_name_calc, axis_name_calc, style)
				}
			}
		}
		index_dif_calc ++
		index_dif_expect = index_dif_calc
		fmt.Println( )
	}

	f_diff.SetActiveSheet(index)
	if err := f_diff.SaveAs("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.compare.xlsx"); err != nil {
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

func GetCols (path  string)  {
	/// 计算值
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	cols, err := f.Cols("ats_result")
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

