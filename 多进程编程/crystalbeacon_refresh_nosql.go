package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"golang.org/x/sys/windows/registry"
	//"github.com/go-redis/redis/v8"
)

func main() {

	redis_info := GetRedisConfigFormRegEdit()

	//Addr:=    "192.168.52.128:46379"
	//Password:= "sparksubmit666"  // no password set
	//DB:=      2                // use default DB
	Remove_keys_via_lua_script(redis_info.Addr, redis_info.Password, redis_info.Dbnum)

}

type RedisInfo struct {
	Addr     string
	Password string
	Dbnum    int
}

func GetRedisConfigFormRegEdit() *RedisInfo {
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
	info := &RedisInfo{
		Addr:     Addr,
		Password: Password,
		Dbnum:    dbnum,
	}
	return info
}

func Remove_keys_via_lua_script(Addr string, Password string, Dbnum int) {
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

///1638278810.894586 [0 192.168.52.1:1374] "UNSUBSCRIBE" "\xbeE\xf0t\xbd+\x92N\xb4b#\xcfqlmo"
//1638278819.665184 [0 192.168.52.1:1373] "INFO" "replication"
//1638278867.305334 [0 192.168.52.1:2335] "auth" "sparksubmit666"
//1638278867.305395 [0 192.168.52.1:2335] "select" "2"
//1638278867.305969 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "monitor:atsinspect*"
//1638278867.306058 [2 lua] "KEYS" "monitor:atsinspect*"
//1638278867.306557 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "ats:apply:*"
//1638278867.306610 [2 lua] "KEYS" "ats:apply:*"
//1638278867.307016 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "ats:class:*"
//1638278867.307072 [2 lua] "KEYS" "ats:class:*"
//1638278867.307735 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "ats:inspect:*"
//1638278867.307812 [2 lua] "KEYS" "ats:inspect:*"
//1638278867.308186 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "ats:inspect_seal:*"
//1638278867.308250 [2 lua] "KEYS" "ats:inspect_seal:*"
//1638278867.308772 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "ats:manual:*"
//1638278867.308835 [2 lua] "KEYS" "ats:manual:*"
//1638278867.309838 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "ats:shifttype:*"
//1638278867.309883 [2 lua] "KEYS" "ats:shifttype:*"
//1638278867.310194 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "ats:system:*"
//1638278867.310236 [2 lua] "KEYS" "ats:system:*"
//1638278867.310916 [2 192.168.52.1:2335] "evalsha" "e2ce9f7b619313b044f28743c3791dd34ce56aad" "1" "sidecar:atsinspect:*"
//1638278867.310973 [2 lua] "KEYS" "sidecar:atsinspect:*"
