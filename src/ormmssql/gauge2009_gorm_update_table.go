package ormmssql

/// go get -u gorm.io/gorm
/// go get -u gorm.io/driver/sqlite

import (
	//"database/sql"

	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"time"
)

type Ats_result_bak2 struct {
	//gorm.Model
	Key_id           string `sql:"type:nvarchar(36);"`
	Emp_id           string `sql:"type:nvarchar(36);"`
	Ats_date         time.Time
	Set_hours        decimal.Decimal `json:"amount" sql:"type:decimal(18,4);"`
	Real_hours       float32         `sql:"type:decimal(10,2);"`
	Shift_class_type string          `sql:"type:nvarchar(100);"`
	Shift_type       string          `sql:"type:nvarchar(100);"`
	Create_dt        time.Time
	Create_by        string `sql:"type:nvarchar(36);"`
	Last_updated_dt  time.Time
	Last_updated_by  string `sql:"type:nvarchar(36);"`
	Ot_policy        string `sql:"type:nvarchar(50);"`
	Base_policy      string `sql:"type:nvarchar(50);"`
	Calendar_id      string `sql:"type:nvarchar(36);"`
	Read_dt          string `sql:"type:nvarchar(36);"`
	Is_post          int
	Unit             string `sql:"type:nvarchar(20);"`
	Inspect_status   string `sql:"type:nvarchar(20);"`
	Inspect_dt       time.Time
	Push_dt          time.Time
	Inspect_batch_id string `sql:"type:nvarchar(36);"`
	Ats_seal_status  string `sql:"type:nvarchar(30);"`
}

/// 实体类表名称默认是ats_results复数形式
type Ats_result struct {
	//gorm.Model
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
	Read_dt          string          `gorm:"size:36"`
	Is_post          bool            `gorm:"type:bit"`
	Unit             string          `gorm:"size:20"`
	Inspect_status   string          `gorm:"size:20"`
	Inspect_dt       time.Time       `gorm:"type:datetime"`
	Push_dt          time.Time       `gorm:"type:datetime"`
	Inspect_batch_id string          `gorm:"size:36"`
	Ats_seal_status  string          `gorm:"size:30"` //`json:"username"  gorm:"index;column:username"`
}

/// 实现TableName接口自定义表名。否则表名称默认是ats_results复数形式
func (this *Ats_result) TableName() string {
	return "ats_result"
}

func Update_Ats_result() {

	decimal.DivisionPrecision = 4 // 保留4位小数，如有更多位，则进行四舍五入保留两位小数

	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://sa:sparksubmit666@192.168.1.7/hive?database=base_api"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Ats_result{})
	var num1 float64 = 13.65432
	var num2 int = 2
	sh := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(float64(num2)))
	rh := decimal.NewFromFloat(num1)
	fmt.Println(sh)
	fmt.Println(rh)

	//tin, err := time.Parse(time.RFC3339, "2006-01-02T22:04:05.787-07:00")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var timeCol civil.Time = civil.TimeOf(tin)
	//var dateCol civil.Date = civil.DateOf(tin)
	//var smalldatetimeCol string = "2006-01-02 22:04:00"
	//var datetimeCol mssql.DateTime1 = mssql.DateTime1(tin)
	//var datetime2Col civil.DateTime = civil.DateTimeOf(tin)
	//var datetimeoffsetCol mssql.DateTimeOffset = mssql.DateTimeOffset(tin)

	t, _ := time.Parse(time.RFC3339, "2021-08-13T17:43:05+07:00")
	//t, _ := time.Parse("2006-01-02 15:04:05", "2021-08-13T17:43:05+07:00")
	fmt.Println(t)
	tt := t.Format("2006-01-02 15:04:05")
	fmt.Println(tt)

	now := time.Now().Format("2006-01-02 15:04:05") //go语言的诞生时间
	fmt.Println(now)

	t2, err := time.ParseInLocation("2006-01-02 15:04:05", now, time.Local)

	// Create
	db.Create(&Ats_result{Key_id: "ats_result_08131626_01", Emp_id: "11-22-33-44-55",
		Ats_date: t2, Create_dt: t2, Inspect_dt: t2, Last_updated_dt: t2, Push_dt: t2,
		Set_hours: sh, Real_hours: rh,
		Is_post: true})
	//db.Create(&Ats_result{Key_id: "ats_result_08131626_01", Emp_id: "11-22-33-44-55",Ats_date: t, Create_dt: time.Now(),Inspect_dt: t })
	//db.Create(&Ats_result{Key_id: "ats_result_08131626_02", Emp_id: "11-22-33-44-66" })

	// Read
	var ats_result Ats_result
	//db.First(&ats_result, "ats_result_08131626_01")
	//db.First(&ats_result, "emp_id = ?", "11-22-33-44-55")
	db.First(&ats_result, "create_dt > ?", "2021-08-16 14:00:00.000")

	var ats_result_2 Ats_result
	db.Where(" set_hours >  ? AND   ats_date>= ?   ORDER BY ats_date DESC", 12, "2021-08-16 14:58:59.000").Find(&ats_result_2)
	fmt.Printf("█ █ █ █ █  value is %v, type is %T\n", ats_result_2.Ats_date, ats_result_2.Ats_date)

	// Inline Condition
	var ats_result_list1 []Ats_result
	//db.Find(&ats_result_list1, " set_hours >  ? AND   ats_date>= ?   ORDER BY ats_date DESC",12, "2021-08-16 16:35:59.000")
	//db.Select("key_id","ats_date","emp_id","set_hours","real_hours","is_post").Find(&ats_result_list1, " set_hours >  ? AND   ats_date>= ?   ORDER BY ats_date DESC",12, "2021-08-16 16:35:59.000")
	//db.Select("key_id","ats_date","emp_id","set_hours","real_hours","is_post").Find(&ats_result_list1, " set_hours >  ? AND   ats_date>= ? ",12, "2021-08-16 16:35:59.000").Order("ats_date desc")
	db.Limit(2).Select("key_id", "ats_date", "emp_id", "set_hours", "real_hours", "is_post").Find(&ats_result_list1, " set_hours >  ? AND   ats_date>= ? ", 12, "2021-08-16 16:35:59.000").Order("ats_date desc")

	fmt.Printf("█ █ █ █ █  ats_result_list count = %v\n", len(ats_result_list1))
	//for   a_result   , _ := range ats_result_list1{
	//	fmt.Printf("a_result 结果:%v\n", (a_result))
	//}
	for i := 0; i < len(ats_result_list1); i++ {
		if i < 4 {
			var at = ats_result_list1[i]
			fmt.Println("█ ", at, " █")
			fmt.Printf("█ █ █ █ █ █ █ █ █ █  a_result 结果:Key_id=%v █ Ats_date=%v █ Emp_id=%v █ Set_hours=%v █ Real_hours=%v █ Is_post=%v\n", at.Key_id, at.Ats_date, at.Emp_id, at.Set_hours, at.Real_hours, at.Is_post)
			//{ats_result_08131626_01 11-22-33-44-55 2021-08-16 17:05:08 +0000 UTC 15.6543 13.6543   2021-08-16 17:05:08 +0000 UTC  2021-08-16 17:05:08 +0000 UTC      true   2021-08-16 17:05:08 +0000 UTC 2021-08-16 17:05:08 +0000 UTC  }
			//{ats_result_08131626_01 11-22-33-44-55 2021-08-16 17:00:45 +0000 UTC 15.6543 13.6543   2021-08-16 17:00:45 +0000 UTC  2021-08-16 17:00:45 +0000 UTC      true   2021-08-16 17:00:45 +0000 UTC 2021-08-16 17:00:45 +0000 UTC  }
			// ......
			//{ats_result_08131626_01 11-22-33-44-55 2021-08-16 16:43:31 +0000 UTC 15.6543 13.6543   2021-08-16 16:43:31 +0000 UTC  2021-08-16 16:43:31 +0000 UTC      true   2021-08-16 16:43:31 +0000 UTC 2021-08-16 16:43:31 +0000 UTC  }
			//{ats_result_08131626_01 11-22-33-44-55 2021-08-16 16:42:39 +0000 UTC 15.6543 13.6543   2021-08-16 16:42:39 +0000 UTC  2021-08-16 16:42:39 +0000 UTC      true   2021-08-16 16:42:39 +0000 UTC 2021-08-16 16:42:39 +0000 UTC  }
		}
	}

	//ats_result_list := list.New()
	//db.Find(&ats_result_list, " set_hours >  ? AND   ats_date>= ?   ORDER BY ats_date DESC",12, "2021-08-16 16:35:59.000")
	//fmt.Printf("█ █ █ █ █  ats_result_list count = %v\n", ats_result_list.Len())
	//for l := ats_result_list.Front(); l != nil; l = l.Next(){
	//	//fmt.Print(l.Value, " ")
	//	fmt.Printf("a_result 结果:%v\n", l.Value)
	//}

	/// Update with conditions and model value

	//db.Model(&ats_result).Update("Ats_seal_status", "N/A")
	//db.Model(&ats_result).Updates(Ats_result{Emp_id: "11-22-33-44-66", Ats_seal_status: "N/A"}) // non-zero fields
	//db.Model(&ats_result).Updates(map[string]interface{}{"Emp_id": "11-22-33-44-66", "Ats_seal_status": "N/A"})
	//db.Model(&ats_result).Updates(map[string]interface{}{"Emp_id": "11-22-33-44-66", "Ats_seal_status": "N/A"})

	db.Model(&ats_result).Where("create_dt > ?", "2021-08-16 14:00:00.000").Update("Ats_seal_status", "N/A")

	probe := db.Model(&ats_result).Where("create_dt > ?", "2021-08-04 17:00:00.000").Updates(map[string]interface{}{"Emp_id": "11-22-33-44-66", "Ats_seal_status": "success"})
	fmt.Printf("█ █ █ █ ██ █ █ █ █   █ █ █ █ ██ █ █ █ █ probe= %v\n", probe.RowsAffected)

	//
	//
	//db.Delete(&ats_result, "ats_result_08131626_01")

}
