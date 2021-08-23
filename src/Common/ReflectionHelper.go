package Common

import (
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
	"time"
)

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
