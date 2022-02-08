package main

import (
	"golang.org/x/sys/windows/registry"
)

func main() {
	// 创建：指定路径的项
	// 路径：HKEY_CURRENT_USER\Software\Hello Go
	key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\ctmw\crystalbeacon`, registry.ALL_ACCESS)
	defer key.Close()

	// 判断是否已经存在了
	if exists {
		println(`键已存在`)
	} else {
		println(`新建注册表键`)
	}
	//Addr:=    "192.168.52.128:46379"
	//Password:= "sparksubmit666"  // no password set
	//DB:=      2                // use default DB
	// 写入：32位整形值
	key.SetDWordValue(`DB`, uint32(2))
	// 写入：64位整形值
	//key.SetQWordValue(`64位整形值`, uint64(123456))
	// 写入：字符串
	key.SetStringValue(`Addr`, `192.168.52.128:46379`)
	key.SetStringValue(`Password`, `sparksubmit666`)
	// 写入：字符串数组
	//key.SetStringsValue(`字符串数组`, []string{`hello`, `world`})
	// 写入：二进制
	//key.SetBinaryValue(`二进制`, []byte{0x11, 0x22})

	// 读取：字符串
	Addr, _, _ := key.GetStringValue(`Addr`)
	println(Addr)
	// 读取：字符串
	Password, _, _ := key.GetStringValue(`Password`)
	println(Password)
	// 读取：字符串
	DB, _, _ := key.GetIntegerValue(`DB`)
	println(DB)

	// 读取：一个项下的所有子项
	keys, _ := key.ReadSubKeyNames(0)
	for _, key_subkey := range keys {
		// 输出所有子项的名字
		println(key_subkey)
	}

	// 创建：子项
	subkey, _, _ := registry.CreateKey(key, `子项`, registry.ALL_ACCESS)
	defer subkey.Close()

	// 删除：子项
	// 该键有子项，所以会删除失败
	// 没有子项，删除成功
	registry.DeleteKey(key, `子项`)
}
