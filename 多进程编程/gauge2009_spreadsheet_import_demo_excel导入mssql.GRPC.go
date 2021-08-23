package main

import (
	"grpclab"
)

// This example shows the usage of Connector type
func main() {
	/// 将excel导入sql server 数据库
	var lab = grpclab.SpreadsheetService{}
	lab.Invoke()
}
