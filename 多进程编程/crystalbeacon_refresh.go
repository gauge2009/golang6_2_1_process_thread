package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"golang.org/x/sys/windows/registry"
)

/// crystalbeacon_refresh
func main() {

	redis_info := GetRedisConfigInfoFormRegEdit()

	//Addr:=    "192.168.52.128:46379"
	//Password:= "sparksubmit666"  // no password set
	//DB:=      2                // use default DB
	Remove_keys_via_lua_script_V2(redis_info.Addr, redis_info.Password, redis_info.Dbnum)
	fmt.Printf("done %+v", redis_info.Addr)
}

type RedisInfoConnInfo struct {
	Addr     string
	Password string
	Dbnum    int
}

func Remove_keys_via_lua_script_V2(Addr string, Password string, Dbnum int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password, // no password set
		DB:       Dbnum,    // use default DB
	})

	Remove_Keys := redis.NewScript(`
		local ks = redis.call('KEYS', KEYS[1])     	 
		for i=1,#ks,5000 do             
			redis.call('del', unpack(ks, i, math.min(i+4999, #ks)) )       
		end 
		return true 
	`)
	var ctxr = context.Background()

	//key_name:= "monitor:atsinspect:0083f73c-cb1d-428b-b788-7ff7008d8d96:Porccess_NeedToCalcCount"
	n, err := Remove_Keys.Run(ctxr, rdb, []string{"monitor:atsinspect*"}).Result()
	n, err = Remove_Keys.Run(ctxr, rdb, []string{"ats:apply:*"}).Result()
	n, err = Remove_Keys.Run(ctxr, rdb, []string{"ats:class:*"}).Result()
	n, err = Remove_Keys.Run(ctxr, rdb, []string{"ats:inspect:*"}).Result()
	n, err = Remove_Keys.Run(ctxr, rdb, []string{"ats:inspect_seal:*"}).Result()
	n, err = Remove_Keys.Run(ctxr, rdb, []string{"ats:manual:*"}).Result()
	n, err = Remove_Keys.Run(ctxr, rdb, []string{"ats:shifttype:*"}).Result()
	n, err = Remove_Keys.Run(ctxr, rdb, []string{"ats:system:*"}).Result()
	n, err = Remove_Keys.Run(ctxr, rdb, []string{"sidecar:atsinspect:*"}).Result()
	fmt.Println(n, err)

}
func GetRedisConfigInfoFormRegEdit() *RedisInfoConnInfo {
	//ExampleClient()
	//Lua_script()
	key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\ctmw\crystalbeacon`, registry.ALL_ACCESS)
	defer key.Close()
	// 判断是否已经存在了
	if exists {
		println(`键已存在`)
	} else {
		println(`新建注册表键`)
	}
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
	//size := binary.BigEndian.Uint32(b[4:])

	DB, _, _ := key.GetIntegerValue(`DB`)
	//n, err := rdr.Discard(int(size))
	dbnum := int(DB)
	println(DB)
	info := &RedisInfoConnInfo{
		Addr:     Addr,
		Password: Password,
		Dbnum:    dbnum,
	}
	return info
}
