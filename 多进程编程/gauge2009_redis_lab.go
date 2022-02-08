package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	//"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	//ExampleClient()
	//Lua_script()
	Remove_keys_by_lua_script()
	Redis_client_string_demo()
	Redis_client_hash_demo()
	Scanning_hash_fields_into_a_struct()
}

type Model struct {
	Str1    string   `redis:"str1"`
	Str2    string   `redis:"str2"`
	Int     int      `redis:"int"`
	Bool    bool     `redis:"bool"`
	Ignored struct{} `redis:"-"`
}

func Scanning_hash_fields_into_a_struct() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.52.128:46379",
		Password: "sparksubmit666", // no password set
		DB:       2,                // use default DB
	})
	/// Because go-redis does not provide a helper to save structs in Redis, we are using a pipeline to load some data into our database:
	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, "hashkey", "str1", "hello")
		rdb.HSet(ctx, "hashkey", "str2", "world")
		rdb.HSet(ctx, "hashkey", "int", 123)
		rdb.HSet(ctx, "hashkey", "bool", 1)
		return nil
	}); err != nil {
		panic(err)
	}
	///After that we are ready to scan the data using HGetAll:
	var model1 Model
	// Scan all fields into the model.
	if err := rdb.HGetAll(ctx, "hashkey").Scan(&model1); err != nil {
		panic(err)
	}
	///Or HMGet:
	var model2 Model
	// Scan a subset of the fields.
	if err := rdb.HMGet(ctx, "hashkey", "str1", "int").Scan(&model2); err != nil {
		panic(err)
	}
	fmt.Printf("model2 str1 = %v\n", model2.Str1)
	fmt.Printf("model2 int = %v\n", model2.Int)

}
func Redis_client_hash_demo() {

}
func Redis_client_string_demo() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.52.128:46379",
		Password: "sparksubmit666", // no password set
		DB:       2,                // use default DB
	})

	err := rdb.Set(ctx, "firstname", "gauge2009", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "firstname").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("firstname", val)

	val2, err := rdb.Get(ctx, "lastname").Result()
	if err == redis.Nil {
		fmt.Println("lastname does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("lastname", val2)
	}
	// Output: key value
	// key2 does not exist
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
func Remove_keys_by_lua_script() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.52.128:46379",
		Password: "sparksubmit666", // no password set
		DB:       2,                // use default DB
	})

	Remove_Keys := redis.NewScript(`
		local ks = redis.call('KEYS', KEYS[1])     	 
		for i=1,#ks,5000 do             
			redis.call('del', unpack(ks, i, math.min(i+4999, #ks)) )       
		end 
		return true 
	`)

	//key_name:= "monitor:atsinspect:0083f73c-cb1d-428b-b788-7ff7008d8d96:Porccess_NeedToCalcCount"
	n, err := Remove_Keys.Run(ctx, rdb, []string{"monitor:atsinspect*"}).Result()
	n, err = Remove_Keys.Run(ctx, rdb, []string{"ats:apply:*"}).Result()
	n, err = Remove_Keys.Run(ctx, rdb, []string{"ats:class:*"}).Result()
	n, err = Remove_Keys.Run(ctx, rdb, []string{"ats:inspect:*"}).Result()
	n, err = Remove_Keys.Run(ctx, rdb, []string{"ats:inspect_seal:*"}).Result()
	n, err = Remove_Keys.Run(ctx, rdb, []string{"ats:manual:*"}).Result()
	n, err = Remove_Keys.Run(ctx, rdb, []string{"ats:shifttype:*"}).Result()
	n, err = Remove_Keys.Run(ctx, rdb, []string{"ats:system:*"}).Result()
	n, err = Remove_Keys.Run(ctx, rdb, []string{"sidecar:atsinspect:*"}).Result()
	fmt.Println(n, err)

}

/// 示例： 模糊删除指定的keys
func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.52.128:46379",
		Password: "sparksubmit666", // no password set
		DB:       2,                // use default DB
	})

	//opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	//if err != nil {
	//	panic(err)
	//}
	//

	key_name := "monitor:atsinspect:0083f73c-cb1d-428b-b788-7ff7008d8d96:Porccess_NeedToCalcCount"
	err := rdb.Set(ctx, key_name, "68823", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, key_name).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key_name, val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func Lua_script() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.52.128:46379",
		Password: "sparksubmit666", // no password set
		DB:       2,                // use default DB
	})

	//key_name:= "monitor:atsinspect:0083f73c-cb1d-428b-b788-7ff7008d8d96:Porccess_NeedToCalcCount"

	IncrByXX := redis.NewScript(`
		if redis.call("GET", KEYS[1]) ~= false then
			return redis.call("INCRBY", KEYS[1], ARGV[1])
		end
		return false
	`)

	n, err := IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

	err = rdb.Set(ctx, "xx_counter", "40", 0).Err()
	if err != nil {
		panic(err)
	}

	n, err = IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

}
