package grpclab

import (
	"github.com/shopspring/decimal"
	"time"
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
