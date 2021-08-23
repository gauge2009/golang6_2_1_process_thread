package spreadsheet

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	//"github.com/xuri/excelize/v2"
	"strconv"
)

//https://pkg.go.dev/fmt#Printf

/// 实体类表名称默认是ats_results复数形式
type Ats_result struct {
	Key_id           string          `gorm:"size:36"`
	Emp_id           string          `gorm:"size:36"`
	Ats_date         time.Time       `gorm:"type:datetime"`
	Set_hours        decimal.Decimal `gorm:"type:decimal(18,4)"`
	Real_hours       decimal.Decimal `gorm:"type:decimal(18,4)"`
	Shift_class_type string          `gorm:"size:100"`
	Shift_type       string          `gorm:"size:100"`
	Create_dt        time.Time       `gorm:"type:datetime"`
	Create_by        string          `gorm:"size:36"`
	Last_updated_dt  time.Time       `gorm:"type:datetime"`
	Last_updated_by  string          `gorm:"size:36"`
	Ot_policy        string          `gorm:"size:50"`
	Base_policy      string          `gorm:"size:50"`
	Calendar_id      string          `gorm:"size:36"`
	Read_dt          time.Time       `gorm:"type:datetime"`
	Is_post          bool            `gorm:"type:bit"`
	Unit             string          `gorm:"size:20"`
	Inspect_status   string          `gorm:"size:20"`
	Inspect_dt       time.Time       `gorm:"type:datetime"`
	Push_dt          time.Time       `gorm:"type:datetime"`
	Inspect_batch_id string          `gorm:"size:36"`
	Ats_seal_status  string          `gorm:"size:30"` //`json:"username"  gorm:"index;column:username"`
	//gorm.Model
}

/// 实现TableName接口自定义表名。否则表名称默认是ats_results复数形式
func (this *Ats_result) TableName() string {
	return "ats_result"
}

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

func Invoke() {
	/// 计算值
	//SmartCheckCalc("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.xlsx")

	//
	///// 预期值
	//SmartCheckCalc("D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.expect.xlsx")

	//path_root := getCurrentAbPathByExecutable() // go run 与 go build 执行不一样，前者使用C:\Users\Administrator\AppData\Local\Temp
	path_root := getCurrentAbPath() // // go run 与 go build  统一使用go run 制指定的
	fmt.Println("getCurrentAbPath = ", path_root)

	sheet_name := "ats_result"
	//
	///// 获得行列数
	//path_calc :="D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.xlsx"
	//path_expect :="D:\\T2_21\\Golang十四章经\\6.2.go高并发分布式与微服务\\6.2.go高并发分布式与微服务（以上后台素材以全）\\1.进程线程复习\\code\\多进程编程\\inspect_data.expect.xlsx"
	path_calc := path_root + "\\golang from excel to sqlserver.xlsx"
	//path_expect := path_root + "\\inspect_data.expect.xlsx"

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

	FromExcelToMssql(path_calc, path_calc, sheet_name, len_rows, len_cell, column_dictionary)

	//SmartGetCols(path_calc,sheet_name)
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

func FromExcelToMssql(path_calc string, path_expect string, sheet_name string, len_rows int, len_cell int, column_dictionary map[int]string) {
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
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}

	//fields := []string{}// 切片
	fieldsMap := make(map[string]string)

	/// 持久化
	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://sa:sparksubmit666@192.168.1.7/hive?database=base_api"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Ats_result{})

	for i := 0; i < len_rows; i++ {
		//if(i>0) {
		ats_result_obj := &Ats_result{
			Key_id:          "123",
			Ats_seal_status: "success",
		}
		valueOfAtsResult := reflect.ValueOf(ats_result_obj).Elem()

		//}
		for j := 0; j < len_cell; j++ {
			axis_name := column_dictionary[j] + strconv.Itoa(i+1)
			prefix := column_dictionary[j]
			field_name := fieldsMap[prefix]
			///先形成表头
			if i == 0 {
				filed_name := rows[i][j]
				fmt.Printf("row[%v], axis_name=%v, value = %v\n", i, axis_name, filed_name)
				//fields = append(fields,filed_name)// 切片后加一个元素
				// 添加映射关系
				fieldsMap[prefix] = filed_name
				continue
			} else {
				cell := rows[i][j]
				/*if(field_name == "key_id") {
					 //property_name := strings.ToLower(field_name)
					 property_name :=  Capitalize(field_name)
					 valueOfFiled := valueOfAtsResult.FieldByName(property_name)
					 // 判断字段的 Value 是否可以设定变量值
					  if valueOfFiled.CanSet() {
					     val := reflect.ValueOf(cell)
						 valueOfFiled.Set(val)
					  }else {
					  	print("oops")
					  }

					 fmt.Printf("AtsResult`s field_name is %s", ats_rslt.Key_id)
				 }*/

				/// █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
				/// █ █ █ █ █ █ reflection  █ █ █ █ █ █
				/// █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
				/// SetValueByProperty(valueOfAtsResult, field_name, cell)
				///fmt.Printf("\nAtsResult`s field_name is %s\n", ats_result_obj.Key_id)
				typeOfAts_result := reflect.TypeOf(Ats_result{})
				//	fmt.Printf("Ats_result's type is %s, kind is %s", typeOfAts_result, typeOfAts_result.Kind())
				SmartSetValueByProperty(typeOfAts_result, valueOfAtsResult, field_name, cell)

				/// 记录预期值
				/*
					fmt.Printf("row[%d], ", i)
					fmt.Printf("axis_name=%s,", axis_name)
					fmt.Printf("value = %v\n", cell)
				*/

			}
		}
		fmt.Printf("================ end of row %v==================\n", i)

		//tn, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
		//db.Create(&Ats_result{Key_id: "ats_result_08131626_01", Emp_id: "11-22-33-44-55",
		//	Ats_date: t2, Create_dt: t2, Inspect_dt: t2, Last_updated_dt: t2, Push_dt: t2,
		//	Set_hours: sh, Real_hours: rh,
		//	Is_post: true})

		if i > 0 {
			//db.Create(&Ats_result{Key_id: ats_result_obj.Key_id, Emp_id: ats_result_obj.Emp_id,
			//	Ats_date: ats_result_obj.Ats_date, Create_dt: ats_result_obj.Create_dt, Inspect_dt: ats_result_obj.Inspect_dt, Last_updated_dt: ats_result_obj.Last_updated_dt, Push_dt: tn,
			//	Set_hours: ats_result_obj.Set_hours, Real_hours: ats_result_obj.Real_hours,
			//	Is_post: ats_result_obj.Is_post})
			//db.Create( ats_result_obj)
			db.Create(&Ats_result{Key_id: ats_result_obj.Key_id, Emp_id: ats_result_obj.Emp_id,
				Ats_date: ats_result_obj.Ats_date, Create_dt: ats_result_obj.Create_dt, Inspect_dt: ats_result_obj.Inspect_dt, Read_dt: ats_result_obj.Read_dt, Last_updated_dt: ats_result_obj.Last_updated_dt, Push_dt: ats_result_obj.Push_dt,
				Set_hours: ats_result_obj.Set_hours, Real_hours: ats_result_obj.Real_hours,
				Is_post:    true,
				Shift_type: ats_result_obj.Shift_type, Shift_class_type: ats_result_obj.Shift_class_type, Unit: ats_result_obj.Unit, Create_by: ats_result_obj.Create_by,
				Ot_policy: ats_result_obj.Ot_policy, Last_updated_by: ats_result_obj.Last_updated_by, Base_policy: ats_result_obj.Base_policy, Calendar_id: ats_result_obj.Calendar_id,
				Inspect_status: ats_result_obj.Inspect_status, Inspect_batch_id: ats_result_obj.Inspect_batch_id, Ats_seal_status: ats_result_obj.Ats_seal_status,
			})
		}

	}
	for _, field := range fieldsMap {
		fmt.Print(field, "\t")
	}

	fmt.Println(" █ █ █ █ █ Create success  █ █ █ █ █")
}

/// 反射，根据类型信息设置实例的某个属性
func SmartSetValueByProperty(typeofClass reflect.Type, valueOfObject reflect.Value, field_name string, cell_val string) {

	fieldsTypeMap := make(map[string]string)
	//// 通过 #NumField 获取结构体字段的数量
	for i := 0; i < typeofClass.NumField(); i++ {
		fn := typeofClass.Field(i).Name
		ft := typeofClass.Field(i).Type
		//fk:=typeofClass.Field(i).Type.Kind()
		//fmt.Printf("field' name is %s, type is %s, kind is %s\n",
		//	fn,
		//	ft,
		//	fk)
		fieldsTypeMap[fn] = ft.String()
	}

	fmt.Printf("█ █ █ █ █  █ █ █ █ █  █ █ █ █ █ %v", field_name)
	property_name := strings.ToLower(field_name)
	property_name = Capitalize(property_name)
	//if field_name == "ats_date" || field_name == "create_dt" || field_name == "read_dt" || field_name == "inspect_dt" || field_name == "push_dt" || field_name == "last_updated_dt" {
	if fieldsTypeMap[property_name] == "time.Time" {
		//now := time.Now().Format("2006-01-02 15:04:05") //go语言的诞生时间
		//fmt.Println(now)
		println("█ █ █ █ █ datetime █ █ █ █ █ ")
		var cell_val_time time.Time
		if cell_val == "" {
			//now := time.Now().Format("2006-01-02 15:04:05") //go语言的诞生时间
			cell_val_time, _ = time.ParseInLocation("2006-01-02 15:04:05", "2111-11-11 11:11:11", time.Local)
		} else {
			cell_val_time, _ = time.ParseInLocation("2006/01/02 15:04:05", cell_val, time.Local)
			if cell_val_time.String() == "0001-01-01 00:00:00 +0000 UTC" {
				cell_val_time, _ = time.ParseInLocation("2006-01-02 15:04:05", "2211-11-11 11:11:11", time.Local)
			}
		}
		println(cell_val_time.String())
		valueOfFiled := valueOfObject.FieldByName(property_name)
		// 判断字段的 Value 是否可以设定变量值
		if valueOfFiled.CanSet() {
			val := reflect.ValueOf(cell_val_time)
			valueOfFiled.Set(val)
		} else {
			print("oops")
		}

	} else if fieldsTypeMap[property_name] == "decimal.Decimal" { //field_name == "set_hours" || field_name == "real_hours" {
		println("█ █ █ █ █ decimal █ █ █ █ █ ")
		cell_val_num, _ := decimal.NewFromString(cell_val)
		valueOfFiled := valueOfObject.FieldByName(property_name)
		// 判断字段的 Value 是否可以设定变量值
		if valueOfFiled.CanSet() {
			val := reflect.ValueOf(cell_val_num)
			valueOfFiled.Set(val)
		} else {
			print("oops")
		}
	} else if fieldsTypeMap[property_name] == "bool" { //field_name == "is_post" {
		println("█ █ █ █ █ bit █ █ █ █ █ ")
		var cell_val_bool bool
		if cell_val == "" || cell_val == "0" {
			//cell_val_bool = nil
			cell_val_bool = false
		} else {
			cell_val_bool = true
		}
		valueOfFiled := valueOfObject.FieldByName(property_name)
		// 判断字段的 Value 是否可以设定变量值
		if valueOfFiled.CanSet() {
			val := reflect.ValueOf(cell_val_bool)
			valueOfFiled.Set(val)
		} else {
			print("oops")
		}
	} else if fieldsTypeMap[property_name] == "string" {
		valueOfFiled := valueOfObject.FieldByName(property_name)
		// 判断字段的 Value 是否可以设定变量值
		if valueOfFiled.CanSet() {
			val := reflect.ValueOf(cell_val)
			valueOfFiled.Set(val)
		} else {
			print("oops")
		}
	}

	fmt.Printf("cell_va = %s", cell_val)
}

/// 反射设置实例的某个属性[缺少类型信息，废弃]
func SetValueByProperty(valueOfObject reflect.Value, field_name string, cell_val string) {
	fmt.Printf("█ █ █ █ █  █ █ █ █ █  █ █ █ █ █ %v", field_name)
	//if cell_val == "" {
	//	return
	//}
	property_name := strings.ToLower(field_name)
	property_name = Capitalize(property_name)
	/*switch field_name {
	case "ats_date":
	case "create_dt":
	case "last_updated_dt":
	case "read_dt":
	case "inspect_dt":
	case "push_dt":
		//now := time.Now().Format("2006-01-02 15:04:05") //go语言的诞生时间
		//fmt.Println(now)
		cell_val, _ := time.ParseInLocation("2006-01-02 15:04:05", cell_val, time.Local)
		println(cell_val.String())

	case "set_hours":
	case "real_hours":

		println("decimal")
	case "is_post":
		println("bit")
	default:
		println("default")
	}*/
	if field_name == "ats_date" || field_name == "create_dt" || field_name == "read_dt" || field_name == "inspect_dt" || field_name == "push_dt" || field_name == "last_updated_dt" {
		//now := time.Now().Format("2006-01-02 15:04:05") //go语言的诞生时间
		//fmt.Println(now)
		println("█ █ █ █ █ datetime █ █ █ █ █ ")
		var cell_val_time time.Time
		if cell_val == "" {
			//now := time.Now().Format("2006-01-02 15:04:05") //go语言的诞生时间
			cell_val_time, _ = time.ParseInLocation("2006-01-02 15:04:05", "2111-11-11 11:11:11", time.Local)
		} else {
			cell_val_time, _ = time.ParseInLocation("2006/01/02 15:04:05", cell_val, time.Local)
			if cell_val_time.String() == "0001-01-01 00:00:00 +0000 UTC" {
				cell_val_time, _ = time.ParseInLocation("2006-01-02 15:04:05", "2211-11-11 11:11:11", time.Local)
			}
		}
		println(cell_val_time.String())
		valueOfFiled := valueOfObject.FieldByName(property_name)
		// 判断字段的 Value 是否可以设定变量值
		if valueOfFiled.CanSet() {
			val := reflect.ValueOf(cell_val_time)
			valueOfFiled.Set(val)
		} else {
			print("oops")
		}

	} else if field_name == "set_hours" || field_name == "real_hours" {
		println("█ █ █ █ █ decimal █ █ █ █ █ ")
		cell_val_num, _ := decimal.NewFromString(cell_val)
		valueOfFiled := valueOfObject.FieldByName(property_name)
		// 判断字段的 Value 是否可以设定变量值
		if valueOfFiled.CanSet() {
			val := reflect.ValueOf(cell_val_num)
			valueOfFiled.Set(val)
		} else {
			print("oops")
		}
	} else if field_name == "is_post" {
		println("█ █ █ █ █ bit █ █ █ █ █ ")
		var cell_val_bool bool
		if cell_val == "" || cell_val == "0" {
			//cell_val_bool = nil
			cell_val_bool = false
		} else {
			cell_val_bool = true
		}
		valueOfFiled := valueOfObject.FieldByName(property_name)
		// 判断字段的 Value 是否可以设定变量值
		if valueOfFiled.CanSet() {
			val := reflect.ValueOf(cell_val_bool)
			valueOfFiled.Set(val)
		} else {
			print("oops")
		}
	} else {
		valueOfFiled := valueOfObject.FieldByName(property_name)
		// 判断字段的 Value 是否可以设定变量值
		if valueOfFiled.CanSet() {
			val := reflect.ValueOf(cell_val)
			valueOfFiled.Set(val)
		} else {
			print("oops")
		}
	}

	fmt.Printf("cell_va = %s", cell_val)
}

///字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
